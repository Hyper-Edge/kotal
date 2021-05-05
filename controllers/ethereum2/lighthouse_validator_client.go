package controllers

import (
	"os"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"
	"github.com/kotalco/kotal/controllers/shared"
)

// LighthouseValidatorClient is SigmaPrime Ethereum 2.0 validator client
type LighthouseValidatorClient struct {
	validator *ethereum2v1alpha1.Validator
}

// Images
const (
	// EnvLighthouseValidatorImage is the environment variable used for SigmaPrime Ethereum 2.0 validator client image
	EnvLighthouseValidatorImage = "LIGHTHOUSE_VALIDATOR_CLIENT_IMAGE"
	// DefaultLighthouseValidatorImage is the default SigmaPrime Ethereum 2.0 validator client image
	DefaultLighthouseValidatorImage = "sigp/lighthouse:v1.3.0"
)

// HomeDir returns container home directory
func (t *LighthouseValidatorClient) HomeDir() string {
	return LighthouseHomeDir
}

// Args returns command line arguments required for client
func (t *LighthouseValidatorClient) Args() (args []string) {

	validator := t.validator

	args = append(args, LighthouseDataDir, shared.PathData(t.HomeDir()))

	args = append(args, LighthouseNetwork, validator.Spec.Network)

	args = append(args, LighthouseDisableAutoDiscover)

	args = append(args, LighthouseInitSlashingProtection)

	if validator.Spec.BeaconEndpoint != "" {
		args = append(args, LighthouseBeaconNodeEndpoint, validator.Spec.BeaconEndpoint)
	}

	if validator.Spec.Graffiti != "" {
		args = append(args, LighthouseGraffiti, validator.Spec.Graffiti)
	}

	return
}

// Command returns command for running the client
func (t *LighthouseValidatorClient) Command() (command []string) {
	command = []string{"lighthouse", "vc"}
	return
}

// Image returns prysm docker image
func (t *LighthouseValidatorClient) Image() string {
	if os.Getenv(EnvLighthouseValidatorImage) == "" {
		return DefaultLighthouseValidatorImage
	}
	return os.Getenv(EnvLighthouseValidatorImage)
}
