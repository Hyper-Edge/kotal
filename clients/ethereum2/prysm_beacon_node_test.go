package ethereum2

import (
	"fmt"
	"os"

	ethereum2v1alpha1 "github.com/kotalco/kotal/apis/ethereum2/v1alpha1"
	sharedAPI "github.com/kotalco/kotal/apis/shared"
	"github.com/kotalco/kotal/controllers/shared"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prysm beacon node", func() {

	node := ethereum2v1alpha1.BeaconNode{
		Spec: ethereum2v1alpha1.BeaconNodeSpec{
			Client:  ethereum2v1alpha1.PrysmClient,
			Network: "mainnet",
		},
	}
	client, _ := NewClient(&node)

	It("Should get correct image", func() {
		// default image
		img := client.Image()
		Expect(img).To(Equal(DefaultPrysmBeaconNodeImage))
		// after changing .spec.image
		testImage := "kotalco/prysm:spec"
		node.Spec.Image = &testImage
		img = client.Image()
		Expect(img).To(Equal(testImage))
		// after setting custom image
		testImage = "kotalco/prysm:test"
		os.Setenv(EnvPrysmBeaconNodeImage, testImage)
		img = client.Image()
		Expect(img).To(Equal(testImage))
	})

	It("Should get correct command", func() {
		Expect(client.Command()).To(ConsistOf("beacon-chain"))
	})

	It("Should get correct env", func() {
		Expect(client.Env()).To(BeNil())
	})

	It("Should get correct home dir", func() {
		Expect(client.HomeDir()).To(Equal(PrysmHomeDir))
	})

	cases := []struct {
		title  string
		node   *ethereum2v1alpha1.BeaconNode
		result []string
	}{
		{
			title: "beacon node syncing mainnet",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:  ethereum2v1alpha1.PrysmClient,
					Network: "mainnet",
					RPC:     true,
					Logging: sharedAPI.WarnLogs,
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				"--mainnet",
				PrysmLogging,
				string(sharedAPI.WarnLogs),
			},
		},
		{
			title: "beacon node syncing mainnet with eth1 endpoint",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:        ethereum2v1alpha1.PrysmClient,
					Network:       "mainnet",
					Eth1Endpoints: []string{"https://localhost:8545"},
					RPC:           true,
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				"--mainnet",
				PrysmWeb3Provider,
				"https://localhost:8545",
			},
		},
		{
			title: "beacon node syncing mainnet with eth1 endpoint and rpc port",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:        ethereum2v1alpha1.PrysmClient,
					Network:       "mainnet",
					Eth1Endpoints: []string{"https://localhost:8545"},
					RPC:           true,
					RPCPort:       9976,
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				"--mainnet",
				PrysmWeb3Provider,
				"https://localhost:8545",
				PrysmRPCPort,
				"9976",
			},
		},
		{
			title: "beacon node syncing mainnet with eth1 endpoint and rpc port and host",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:        ethereum2v1alpha1.PrysmClient,
					Network:       "mainnet",
					Eth1Endpoints: []string{"https://localhost:8545"},
					RPC:           true,
					RPCPort:       9976,
					RPCHost:       "0.0.0.0",
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				"--mainnet",
				PrysmWeb3Provider,
				"https://localhost:8545",
				PrysmRPCPort,
				"9976",
				PrysmRPCHost,
				"0.0.0.0",
			},
		},
		{
			title: "beacon node syncing mainnet with eth1 endpoint and grpc disabled",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:        ethereum2v1alpha1.PrysmClient,
					Network:       "mainnet",
					Eth1Endpoints: []string{"https://localhost:8545"},
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				"--mainnet",
				PrysmWeb3Provider,
				"https://localhost:8545",
				PrysmDisableGRPC,
			},
		},
		{
			title: "beacon node syncing mainnet with eth1 endpoint, certificate and grpc with port",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:  ethereum2v1alpha1.PrysmClient,
					Network: "mainnet",
					Eth1Endpoints: []string{
						"https://localhost:8545",
						"https://localhost:8546",
						"https://localhost:8547",
					},
					GRPC:           true,
					GRPCPort:       4445,
					CertSecretName: "my-certificate",
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				"--mainnet",
				PrysmWeb3Provider,
				"https://localhost:8545",
				PrysmFallbackWeb3Provider,
				"https://localhost:8546",
				PrysmFallbackWeb3Provider,
				"https://localhost:8547",
				PrysmGRPCPort,
				"4445",
				PrysmGRPCGatewayCorsDomains,
				"*",
				PrysmTLSCert,
				fmt.Sprintf("%s/tls.crt", shared.PathSecrets(PrysmHomeDir)),
				PrysmTLSKey,
				fmt.Sprintf("%s/tls.key", shared.PathSecrets(PrysmHomeDir)),
			},
		},
		{
			title: "beacon node syncing mainnet with eth1 endpoint and grpc with port and host",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:  ethereum2v1alpha1.PrysmClient,
					Network: "mainnet",
					Eth1Endpoints: []string{
						"https://localhost:8545",
						"https://localhost:8546",
					},
					GRPC:     true,
					GRPCPort: 4445,
					GRPCHost: "0.0.0.0",
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				"--mainnet",
				PrysmWeb3Provider,
				"https://localhost:8545",
				PrysmFallbackWeb3Provider,
				"https://localhost:8546",
				PrysmGRPCPort,
				"4445",
				PrysmGRPCHost,
				"0.0.0.0",
				PrysmGRPCGatewayCorsDomains,
				"*",
			},
		},
		{
			title: "beacon node syncing mainnet with p2p port, eth1 endpoint and grpc with port and host",
			node: &ethereum2v1alpha1.BeaconNode{
				Spec: ethereum2v1alpha1.BeaconNodeSpec{
					Client:  ethereum2v1alpha1.PrysmClient,
					P2PPort: 7891,
					Network: "mainnet",
					Eth1Endpoints: []string{
						"https://localhost:8545",
						"https://localhost:8546",
					},
					GRPC:     true,
					GRPCPort: 4445,
					GRPCHost: "0.0.0.0",
				},
			},
			result: []string{
				PrysmAcceptTermsOfUse,
				PrysmDataDir,
				PrysmP2PTCPPort,
				"7891",
				PrysmP2PUDPPort,
				"7891",
				"--mainnet",
				PrysmWeb3Provider,
				"https://localhost:8545",
				PrysmFallbackWeb3Provider,
				"https://localhost:8546",
				PrysmGRPCPort,
				"4445",
				PrysmGRPCHost,
				"0.0.0.0",
				PrysmGRPCGatewayCorsDomains,
				"*",
			},
		},
	}

	for _, c := range cases {
		func() {
			cc := c
			It(fmt.Sprintf("Should create correct client arguments for %s", cc.title), func() {
				cc.node.Default()
				client, _ := NewClient(cc.node)
				args := client.Args()
				Expect(args).To(ContainElements(cc.result))
			})
		}()
	}

})
