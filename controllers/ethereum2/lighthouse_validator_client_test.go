package controllers

import (
	"fmt"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lighthouse Ethereum 2.0 validator client arguments", func() {

	cases := []struct {
		title     string
		validator *ethereum2v1alpha1.Validator
		result    []string
	}{
		{
			title: "mainnet validator client",
			validator: &ethereum2v1alpha1.Validator{
				Spec: ethereum2v1alpha1.ValidatorSpec{
					Client:         ethereum2v1alpha1.LighthouseClient,
					Network:        "mainnet",
					BeaconEndpoint: "http://localhost:8899",
					Graffiti:       "Validated by Kotal",
				},
			},
			result: []string{
				LighthouseDataDir,
				LighthouseNetwork,
				"mainnet",
				LighthouseBeaconNodeEndpoint,
				"http://localhost:8899",
				LighthouseGraffiti,
				"Validated by Kotal",
				LighthouseDisableAutoDiscover,
				LighthouseInitSlashingProtection,
			},
		},
	}

	for _, c := range cases {
		func() {
			cc := c
			It(fmt.Sprintf("Should create correct client arguments for %s", cc.title), func() {
				cc.validator.Default()
				client, _ := NewValidatorClient(cc.validator)
				args := client.Args()
				Expect(args).To(ContainElements(cc.result))
			})
		}()
	}

})
