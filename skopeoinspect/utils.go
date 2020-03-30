package skopeoinspect

import (
	"context"
	_ "github.com/containers/image/v5/docker"
	_ "github.com/containers/image/v5/image"
	_ "github.com/containers/image/v5/manifest"
	"github.com/containers/image/v5/pkg/compression"
	_ "github.com/containers/image/v5/transports"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
	_ "github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// optionalBool is a boolean with a separate presence flag.
type optionalBool struct {
	present bool
	value   bool
}

// optionalBool is a cli.Generic == flag.Value implementation equivalent to
// the one underlying flag.Bool, except that it records whether the flag has been set.
// This is distinct from optionalBool to (pretend to) force callers to use
// newOptionalBool
type optionalBoolValue optionalBool

func (ob *optionalBoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	ob.value = v
	ob.present = true
	return nil
}

func (ob *optionalBoolValue) String() string {
	if !ob.present {
		return "" // This is, sadly, not round-trip safe: --flag is interpreted as --flag=true
	}
	return strconv.FormatBool(ob.value)
}

func (ob *optionalBoolValue) IsBoolFlag() bool {
	return true
}

// optionalString is a string with a separate presence flag.
type optionalString struct {
	present bool
	value   string
}

// optionalString is a cli.Generic == flag.Value implementation equivalent to
// the one underlying flag.String, except that it records whether the flag has been set.
// This is distinct from optionalString to (pretend to) force callers to use
// newoptionalString
type optionalStringValue optionalString

func (ob *optionalStringValue) Set(s string) error {
	ob.value = s
	ob.present = true
	return nil
}

func (ob *optionalStringValue) String() string {
	if !ob.present {
		return "" // This is, sadly, not round-trip safe: --flag= is interpreted as {present:true, value:""}
	}
	return ob.value
}

// optionalInt is a int with a separate presence flag.
type optionalInt struct {
	present bool
	value   int
}

// optionalInt is a cli.Generic == flag.Value implementation equivalent to
// the one underlying flag.Int, except that it records whether the flag has been set.
// This is distinct from optionalInt to (pretend to) force callers to use
// newoptionalIntValue
type optionalIntValue optionalInt

func (ob *optionalIntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	ob.value = int(v)
	ob.present = true
	return nil
}

func (ob *optionalIntValue) String() string {
	if !ob.present {
		return "" // If the value is not present, just return an empty string, any other value wouldn't make sense.
	}
	return strconv.Itoa(int(ob.value))
}


// SharedImageOptions collects CLI flags which are image-related, but do not change across images.
// This really should be a part of GlobalOptions, but that would break existing users of (skopeo copy --authfile=).
type SharedImageOptions struct {
	authFilePath string // Path to a */containers/auth.json
}

// ImageOptions collects CLI flags specific to the "docker" transport, which are
// the same across subcommands, but may be different for each image
// (e.g. may differ between the source and destination of a copy)
type DockerImageOptions struct {
	Global         *GlobalOptions      // May be shared across several ImageOptions instances.
	Shared         *SharedImageOptions // May be shared across several ImageOptions instances.
	authFilePath   optionalString      // Path to a */containers/auth.json (prefixed version to override shared image option).
	credsOption    optionalString      // username[:password] for accessing a registry
	dockerCertPath string              // A directory using Docker-like *.{crt,cert,key} files for connecting to a registry or a daemon
	tlsVerify      optionalBool        // Require HTTPS and verify certificates (for docker: and docker-daemon:)
	noCreds        bool                // Access the registry anonymously
}

// ImageOptions collects CLI flags which are the same across subcommands, but may be different for each image
// (e.g. may differ between the source and destination of a copy)
type ImageOptions struct {
	DockerImageOptions
	sharedBlobDir    string // A directory to use for OCI blobs, shared across repositories
	dockerDaemonHost string // docker-daemon: host to connect to
}

// newSystemContext returns a *types.SystemContext corresponding to opts.
// It is guaranteed to return a fresh instance, so it is safe to make additional updates to it.
func (opts *ImageOptions) newSystemContext() (*types.SystemContext, error) {
	ctx := &types.SystemContext{
		RegistriesDirPath:        opts.Global.registriesDirPath,
		ArchitectureChoice:       opts.Global.overrideArch,
		OSChoice:                 opts.Global.overrideOS,
		VariantChoice:            opts.Global.overrideVariant,
		DockerCertPath:           opts.dockerCertPath,
		OCISharedBlobDirPath:     opts.sharedBlobDir,
		AuthFilePath:             opts.Shared.authFilePath,
		DockerDaemonHost:         opts.dockerDaemonHost,
		DockerDaemonCertPath:     opts.dockerCertPath,
		SystemRegistriesConfPath: opts.Global.registriesConfPath,
	}
	if opts.DockerImageOptions.authFilePath.present {
		ctx.AuthFilePath = opts.DockerImageOptions.authFilePath.value
	}
	if opts.tlsVerify.present {
		ctx.DockerDaemonInsecureSkipTLSVerify = !opts.tlsVerify.value
	}
	// DEPRECATED: We support this for backward compatibility, but override it if a per-image flag is provided.
	if opts.Global.tlsVerify.present {
		ctx.DockerInsecureSkipTLSVerify = types.NewOptionalBool(!opts.Global.tlsVerify.value)
	}
	if opts.tlsVerify.present {
		ctx.DockerInsecureSkipTLSVerify = types.NewOptionalBool(!opts.tlsVerify.value)
	}
	if opts.credsOption.present && opts.noCreds {
		return nil, errors.New("creds and no-creds cannot be specified at the same time")
	}
	if opts.credsOption.present {
		var err error
		ctx.DockerAuthConfig, err = getDockerAuth(opts.credsOption.value)
		if err != nil {
			return nil, err
		}
	}
	if opts.noCreds {
		ctx.DockerAuthConfig = &types.DockerAuthConfig{}
	}

	return ctx, nil
}

// imageDestOptions is a superset of ImageOptions specialized for iamge destinations.
type imageDestOptions struct {
	*ImageOptions
	dirForceCompression         bool        // Compress layers when saving to the dir: transport
	ociAcceptUncompressedLayers bool        // Whether to accept uncompressed layers in the oci: transport
	compressionFormat           string      // Format to use for the compression
	compressionLevel            optionalInt // Level to use for the compression
}

// newSystemContext returns a *types.SystemContext corresponding to opts.
// It is guaranteed to return a fresh instance, so it is safe to make additional updates to it.
func (opts *imageDestOptions) newSystemContext() (*types.SystemContext, error) {
	ctx, err := opts.ImageOptions.newSystemContext()
	if err != nil {
		return nil, err
	}

	ctx.DirForceCompress = opts.dirForceCompression
	ctx.OCIAcceptUncompressedLayers = opts.ociAcceptUncompressedLayers
	if opts.compressionFormat != "" {
		cf, err := compression.AlgorithmByName(opts.compressionFormat)
		if err != nil {
			return nil, err
		}
		ctx.CompressionFormat = &cf
	}
	if opts.compressionLevel.present {
		ctx.CompressionLevel = &opts.compressionLevel.value
	}
	return ctx, err
}

func parseCreds(creds string) (string, string, error) {
	if creds == "" {
		return "", "", errors.New("credentials can't be empty")
	}
	up := strings.SplitN(creds, ":", 2)
	if len(up) == 1 {
		return up[0], "", nil
	}
	if up[0] == "" {
		return "", "", errors.New("username can't be empty")
	}
	return up[0], up[1], nil
}

func getDockerAuth(creds string) (*types.DockerAuthConfig, error) {
	username, password, err := parseCreds(creds)
	if err != nil {
		return nil, err
	}
	return &types.DockerAuthConfig{
		Username: username,
		Password: password,
	}, nil
}

// parseImage converts image URL-like string to an initialized handler for that image.
// The caller must call .Close() on the returned ImageCloser.
func parseImage(ctx context.Context, opts *ImageOptions, name string) (types.ImageCloser, error) {
	ref, err := alltransports.ParseImageName(name)
	if err != nil {
		return nil, err
	}
	sys, err := opts.newSystemContext()
	if err != nil {
		return nil, err
	}
	return ref.NewImage(ctx, sys)
}

// parseImageSource converts image URL-like string to an ImageSource.
// The caller must call .Close() on the returned ImageSource.
func parseImageSource(ctx context.Context, opts *ImageOptions, name string) (types.ImageSource, error) {
	ref, err := alltransports.ParseImageName(name)
	if err != nil {
		return nil, err
	}
	sys, err := opts.newSystemContext()
	if err != nil {
		return nil, err
	}
	return ref.NewImageSource(ctx, sys)
}

//var neededCapabilities = []capability.Cap{
//	capability.CAP_CHOWN,
//	capability.CAP_DAC_OVERRIDE,
//	capability.CAP_FOWNER,
//	capability.CAP_FSETID,
//	capability.CAP_MKNOD,
//	capability.CAP_SETFCAP,
//}
//
//func maybeReexec() error {
//	// With Skopeo we need only the subset of the root capabilities necessary
//	// for pulling an image to the storage.  Do not attempt to create a namespace
//	// if we already have the capabilities we need.
//	capabilities, err := capability.NewPid(0)
//	if err != nil {
//		return errors.Wrapf(err, "error reading the current capabilities sets")
//	}
//	for _, cap := range neededCapabilities {
//		if !capabilities.Get(capability.EFFECTIVE, cap) {
//			// We miss a capability we need, create a user namespaces
//			MaybeReexecUsingUserNamespace(true)
//			return nil
//		}
//	}
//	return nil
//}
//
//func reexecIfNecessaryForImages(imageNames ...string) error {
//	// Check if container-storage is used before doing unshare
//	for _, imageName := range imageNames {
//		transport := alltransports.TransportFromImageName(imageName)
//		// Hard-code the storage name to avoid a reference on c/image/storage.
//		// See https://github.com/containers/skopeo/issues/771#issuecomment-563125006.
//		if transport != nil && transport.Name() == "containers-storage" {
//			return maybeReexec()
//		}
//	}
//	return nil
//}
