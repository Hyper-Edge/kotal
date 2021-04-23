package controllers

import (
	"fmt"
	"os"
	"strings"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"
	"github.com/kotalco/kotal/controllers/shared"
)

// TekuValidatorClient is Teku validator client
type TekuValidatorClient struct {
	validator *ethereum2v1alpha1.Validator
}

const (
	// EnvTekuValidatorImage is the environment variable used for PegaSys Teku validator client image
	EnvTekuValidatorImage = "TEKU_VALIDATOR_CLIENT_IMAGE"
	// DefaultTekuValidatorImage is PegaSys Teku validator client image
	DefaultTekuValidatorImage = "consensys/teku:21.2.0"
)

// HomeDir returns container home directory
func (t *TekuValidatorClient) HomeDir() string {
	return TekuHomeDir
}

// Args returns command line arguments required for client
func (t *TekuValidatorClient) Args() (args []string) {

	validator := t.validator

	args = append(args, TekuVC)

	args = append(args, TekuDataPath, shared.PathData(t.HomeDir()))

	args = append(args, TekuNetwork, validator.Spec.Network)

	args = append(args, TekuValidatorsKeystoreLockingEnabled, "false")

	if validator.Spec.BeaconEndpoint != "" {
		args = append(args, TekuBeaconNodeEndpoint, validator.Spec.BeaconEndpoint)
	}

	if validator.Spec.Graffiti != "" {
		args = append(args, TekuGraffiti, validator.Spec.Graffiti)
	}

	keyPass := []string{}
	for i, keystore := range validator.Spec.Keystores {
		path := fmt.Sprintf("%s/validator-keys/%s", shared.PathSecrets(t.HomeDir()), keystore.SecretName)
		keyPass = append(keyPass, fmt.Sprintf("%s/keystore-%d.json:%s/password.txt", path, i, path))
	}

	args = append(args, TekuValidatorKeys, strings.Join(keyPass, ","))

	return args
}

// Command returns command for running the client
func (t *TekuValidatorClient) Command() (command []string) {
	return
}

// Image returns teku docker image
func (t *TekuValidatorClient) Image() string {
	if os.Getenv(EnvTekuValidatorImage) == "" {
		return DefaultTekuValidatorImage
	}
	return os.Getenv(EnvTekuValidatorImage)
}
