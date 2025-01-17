package v1alpha1

// Resources
const (
	// DefaultNodeCPURequest is the cpu requested by Aptos node
	DefaultNodeCPURequest = "2"
	// DefaultNodeCPULimit is the cpu limit for Aptos node
	DefaultNodeCPULimit = "4"

	// DefaultNodeMemoryRequest is the memory requested by Aptos node
	DefaultNodeMemoryRequest = "4Gi"
	// DefaultNodeMemoryLimit is the memory limit for Aptos node
	DefaultNodeMemoryLimit = "8Gi"

	// DefaultNodeStorageRequest is the Storage requested by Aptos node
	DefaultNodeStorageRequest = "250Gi"
)

const (
	// DefaultAPIPort is the default API server port
	DefaultAPIPort uint = 8080
	// DefaultFullnodeP2PPort is the default full node p2p port
	DefaultFullnodeP2PPort uint = 6182
	// DefaultValidatorP2PPort is the default validator node p2p port
	DefaultValidatorP2PPort uint = 6180
	// DefaultHost is the default host
	DefaultHost = "0.0.0.0"
)
