package controllers

import (
	"fmt"
	"os"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"
	"github.com/kotalco/kotal/controllers/shared"
)

// PrysmValidatorClient is Prysmatic labs validator client
type PrysmValidatorClient struct {
	validator *ethereum2v1alpha1.Validator
}

// Images
const (
	// EnvPrysmValidatorImage is the environment variable used for Prysmatic Labs validator client image
	EnvPrysmValidatorImage = "PRYSM_VALIDATOR_CLIENT_IMAGE"
	// DefaultPrysmValidatorImage is Prysmatic Labs validator client image
	DefaultPrysmValidatorImage = "gcr.io/prysmaticlabs/prysm/validator:v1.3.8-hotfix-6c0942"
)

// HomeDir returns container home directory
func (t *PrysmValidatorClient) HomeDir() string {
	return PrysmHomeDir
}

// Args returns command line arguments required for client
func (t *PrysmValidatorClient) Args() (args []string) {

	validator := t.validator

	args = append(args, PrysmAcceptTermsOfUse)

	args = append(args, PrysmDataDir, shared.PathData(t.HomeDir()))

	args = append(args, PrysmWalletDir, fmt.Sprintf("%s/prysm-wallet", shared.PathData(t.HomeDir())))

	args = append(args, PrysmWalletPasswordFile, fmt.Sprintf("%s/prysm-wallet/prysm-wallet-password.txt", shared.PathSecrets(t.HomeDir())))

	args = append(args, fmt.Sprintf("--%s", validator.Spec.Network))

	if validator.Spec.BeaconEndpoint != "" {
		args = append(args, PrysmBeaconRPCProvider, validator.Spec.BeaconEndpoint)
	}

	if validator.Spec.Graffiti != "" {
		args = append(args, PrysmGraffiti, validator.Spec.Graffiti)
	}

	return args
}

// Command returns command for running the client
func (t *PrysmValidatorClient) Command() (command []string) {
	return
}

// Image returns prysm docker image
func (t *PrysmValidatorClient) Image() string {
	if os.Getenv(EnvPrysmValidatorImage) == "" {
		return DefaultPrysmValidatorImage
	}
	return os.Getenv(EnvPrysmValidatorImage)
}
