package skopeoinspect

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/image"
	"github.com/containers/image/v5/manifest"
	"github.com/containers/image/v5/signature"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
	"time"
)

// inspectOutput is the output format of (skopeo inspect), primarily so that we can format it with a simple json.MarshalIndent.
type inspectOutput struct {
	Name          string `json:",omitempty"`
	Tag           string `json:",omitempty"`
	Digest        digest.Digest
	RepoTags      []string
	Created       *time.Time
	DockerVersion string
	Labels        map[string]string
	Architecture  string
	Os            string
	Layers        []string
	Env           []string
}

type GlobalOptions struct {
	debug              bool          // Enable debug output
	tlsVerify          optionalBool  // Require HTTPS and verify certificates (for docker: and docker-daemon:)
	policyPath         string        // Path to a signature verification policy file
	insecurePolicy     bool          // Use an "allow everything" signature verification policy
	registriesDirPath  string        // Path to a "registries.d" registry configuration directory
	overrideArch       string        // Architecture to use for choosing images, instead of the runtime one
	overrideOS         string        // OS to use for choosing images, instead of the runtime one
	overrideVariant    string        // Architecture variant to use for choosing images, instead of the runtime one
	commandTimeout     time.Duration // Timeout for the command execution
	registriesConfPath string        // Path to the "registries.conf" file
}

type InspectOptions struct {
	Global *GlobalOptions
	Image  *ImageOptions
	raw    bool // Output the raw manifest instead of parsing information about the image
	config bool // Output the raw config blob instead of parsing information about the image
}

// getPolicyContext returns a *signature.PolicyContext based on opts.
func (opts *GlobalOptions) getPolicyContext() (*signature.PolicyContext, error) {
	var policy *signature.Policy // This could be cached across calls in opts.
	var err error
	if opts.insecurePolicy {
		policy = &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	} else if opts.policyPath == "" {
		policy, err = signature.DefaultPolicy(nil)
	} else {
		policy, err = signature.NewPolicyFromFile(opts.policyPath)
	}
	if err != nil {
		return nil, err
	}
	return signature.NewPolicyContext(policy)
}

// commandTimeoutContext returns a context.Context and a cancellation callback based on opts.
// The caller should usually "defer cancel()" immediately after calling this.
func (opts *GlobalOptions) commandTimeoutContext() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	var cancel context.CancelFunc = func() {}
	if opts.commandTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, opts.commandTimeout)
	}
	return ctx, cancel
}

func (opts *InspectOptions) Run(args []string, stdout io.Writer) (retErr error) {
	ctx, cancel := opts.Global.commandTimeoutContext()
	defer cancel()

	if len(args) != 1 {
		return errors.New("Exactly one argument expected")
	}
	imageName := args[0]

	//if err := reexecIfNecessaryForImages(imageName); err != nil {
	//	return err
	//}

	sys, err := opts.Image.newSystemContext()
	if err != nil {
		return err
	}

	src, err := parseImageSource(ctx, opts.Image, imageName)
	if err != nil {
		return fmt.Errorf("Error parsing image name %q: %v", imageName, err)
	}

	defer func() {
		if err := src.Close(); err != nil {
			retErr = errors.Wrapf(retErr, fmt.Sprintf("(could not close image: %v) ", err))
		}
	}()

	rawManifest, _, err := src.GetManifest(ctx, nil)
	if err != nil {
		return fmt.Errorf("Error retrieving manifest for image: %v", err)
	}

	if opts.raw && !opts.config {
		_, err := stdout.Write(rawManifest)
		if err != nil {
			return fmt.Errorf("Error writing manifest to standard output: %v", err)
		}
		return nil
	}

	img, err := image.FromUnparsedImage(ctx, sys, image.UnparsedInstance(src, nil))
	if err != nil {
		return fmt.Errorf("Error parsing manifest for image: %v", err)
	}

	if opts.config && opts.raw {
		configBlob, err := img.ConfigBlob(ctx)
		if err != nil {
			return fmt.Errorf("Error reading configuration blob: %v", err)
		}
		_, err = stdout.Write(configBlob)
		if err != nil {
			return fmt.Errorf("Error writing configuration blob to standard output: %v", err)
		}
		return nil
	} else if opts.config {
		config, err := img.OCIConfig(ctx)
		if err != nil {
			return fmt.Errorf("Error reading OCI-formatted configuration data: %v", err)
		}
		err = json.NewEncoder(stdout).Encode(config)
		if err != nil {
			return fmt.Errorf("Error writing OCI-formatted configuration data to standard output: %v", err)
		}
		return nil
	}

	imgInspect, err := img.Inspect(ctx)
	if err != nil {
		return err
	}
	outputData := inspectOutput{
		Name: "", // Set below if DockerReference() is known
		Tag:  imgInspect.Tag,
		// Digest is set below.
		RepoTags:      []string{}, // Possibly overriden for docker.Transport.
		Created:       imgInspect.Created,
		DockerVersion: imgInspect.DockerVersion,
		Labels:        imgInspect.Labels,
		Architecture:  imgInspect.Architecture,
		Os:            imgInspect.Os,
		Layers:        imgInspect.Layers,
		Env:           imgInspect.Env,
	}
	outputData.Digest, err = manifest.Digest(rawManifest)
	if err != nil {
		return fmt.Errorf("Error computing manifest digest: %v", err)
	}
	if dockerRef := img.Reference().DockerReference(); dockerRef != nil {
		outputData.Name = dockerRef.Name()
	}
	if img.Reference().Transport() == docker.Transport {
		sys, err := opts.Image.newSystemContext()
		if err != nil {
			return err
		}
		outputData.RepoTags, err = docker.GetRepositoryTags(ctx, sys, img.Reference())
		if err != nil {
			// some registries may decide to block the "list all tags" endpoint
			// gracefully allow the inspect to continue in this case. Currently
			// the IBM Bluemix container registry has this restriction.
			// In addition, AWS ECR rejects it with 403 (Forbidden) if the "ecr:ListImages"
			// action is not allowed.
			if !strings.Contains(err.Error(), "401") && !strings.Contains(err.Error(), "403") {
				return fmt.Errorf("Error determining repository tags: %v", err)
			}
			logrus.Warnf("Registry disallows tag list retrieval; skipping")
		}
	}
	out, err := json.MarshalIndent(outputData, "", "    ")
	if err != nil {
		return err
	}
	fmt.Fprintf(stdout, "%s\n", string(out))
	return nil
}

