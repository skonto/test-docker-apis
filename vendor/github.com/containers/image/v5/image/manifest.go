package image

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/docker/reference"
	"github.com/containers/image/v5/manifest"
	"github.com/containers/image/v5/types"
	imgspecv1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
)

// genericManifest is an interface for parsing, modifying image manifests and related data.
// Note that the public methods are intended to be a subset of types.Image
// so that embedding a genericManifest into structs works.
// will support v1 one day...
type genericManifest interface {
	serialize() ([]byte, error)
	manifestMIMEType() string
	// ConfigInfo returns a complete BlobInfo for the separate config object, or a BlobInfo{Digest:""} if there isn't a separate object.
	// Note that the config object may not exist in the underlying storage in the return value of UpdatedImage! Use ConfigBlob() below.
	ConfigInfo() types.BlobInfo
	// ConfigBlob returns the blob described by ConfigInfo, iff ConfigInfo().Digest != ""; nil otherwise.
	// The result is cached; it is OK to call this however often you need.
	ConfigBlob(context.Context) ([]byte, error)
	// OCIConfig returns the image configuration as per OCI v1 image-spec. Information about
	// layers in the resulting configuration isn't guaranteed to be returned to due how
	// old image manifests work (docker v2s1 especially).
	OCIConfig(context.Context) (*imgspecv1.Image, error)
	// LayerInfos returns a list of BlobInfos of layers referenced by this image, in order (the root layer first, and then successive layered layers).
	// The Digest field is guaranteed to be provided; Size may be -1.
	// WARNING: The list may contain duplicates, and they are semantically relevant.
	LayerInfos() []types.BlobInfo
	// EmbeddedDockerReferenceConflicts whether a Docker reference embedded in the manifest, if any, conflicts with destination ref.
	// It returns false if the manifest does not embed a Docker reference.
	// (This embedding unfortunately happens for Docker schema1, please do not add support for this in any new formats.)
	EmbeddedDockerReferenceConflicts(ref reference.Named) bool
	// Inspect returns various information for (skopeo inspect) parsed from the manifest and configuration.
	Inspect(context.Context) (*types.ImageInspectInfo, error)
	// UpdatedImageNeedsLayerDiffIDs returns true iff UpdatedImage(options) needs InformationOnly.LayerDiffIDs.
	// This is a horribly specific interface, but computing InformationOnly.LayerDiffIDs can be very expensive to compute
	// (most importantly it forces us to download the full layers even if they are already present at the destination).
	UpdatedImageNeedsLayerDiffIDs(options types.ManifestUpdateOptions) bool
	// UpdatedImage returns a types.Image modified according to options.
	// This does not change the state of the original Image object.
	UpdatedImage(ctx context.Context, options types.ManifestUpdateOptions) (types.Image, error)
	// SupportsEncryption returns if encryption is supported for the manifest type
	//
	// Deprecated: Initially used to determine if a manifest can be copied from a source manifest type since
	// the process of updating a manifest between different manifest types was to update then convert.
	// This resulted in some fields in the update being lost. This has been fixed by: https://github.com/containers/image/pull/836
	SupportsEncryption(ctx context.Context) bool
}

// manifestInstanceFromBlob returns a genericManifest implementation for (manblob, mt) in src.
// If manblob is a manifest list, it implicitly chooses an appropriate image from the list.
func manifestInstanceFromBlob(ctx context.Context, sys *types.SystemContext, src types.ImageSource, manblob []byte, mt string) (genericManifest, error) {
	switch manifest.NormalizedMIMEType(mt) {
	case manifest.DockerV2Schema1MediaType, manifest.DockerV2Schema1SignedMediaType:
		return manifestSchema1FromManifest(manblob)
	case imgspecv1.MediaTypeImageManifest:
		return manifestOCI1FromManifest(src, manblob)
	case manifest.DockerV2Schema2MediaType:
		return manifestSchema2FromManifest(src, manblob)
	case manifest.DockerV2ListMediaType:
		return manifestSchema2FromManifestList(ctx, sys, src, manblob)
	case imgspecv1.MediaTypeImageIndex:
		return manifestOCI1FromImageIndex(ctx, sys, src, manblob)
	default: // Note that this may not be reachable, manifest.NormalizedMIMEType has a default for unknown values.
		return nil, fmt.Errorf("Unimplemented manifest MIME type %s", mt)
	}
}

// manifestLayerInfosToBlobInfos extracts a []types.BlobInfo from a []manifest.LayerInfo.
func manifestLayerInfosToBlobInfos(layers []manifest.LayerInfo) []types.BlobInfo {
	blobs := make([]types.BlobInfo, len(layers))
	for i, layer := range layers {
		blobs[i] = layer.BlobInfo
	}
	return blobs
}

// manifestConvertFn is used to encapsulate helper manifest converstion functions
// to perform applying of manifest update information.
type manifestConvertFn func(context.Context, types.ManifestUpdateInformation) (types.Image, error)

// convertManifestIfRequiredWithUpdate will run conversion functions of a manifest if
// required and re-apply the options to the converted type.
// It returns (nil, nil) if no conversion was requested.
func convertManifestIfRequiredWithUpdate(ctx context.Context, options types.ManifestUpdateOptions, converters map[string]manifestConvertFn) (types.Image, error) {
	if options.ManifestMIMEType == "" {
		return nil, nil
	}

	converter, ok := converters[options.ManifestMIMEType]
	if !ok {
		return nil, errors.Errorf("Unsupported conversion type: %v", options.ManifestMIMEType)
	}

	tmp, err := converter(ctx, options.InformationOnly)
	if err != nil {
		return nil, err
	}

	optionsCopy := options
	optionsCopy.ManifestMIMEType = ""
	return tmp.UpdatedImage(ctx, optionsCopy)
}
