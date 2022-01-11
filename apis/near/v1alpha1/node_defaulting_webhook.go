package v1alpha1

import (
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// +kubebuilder:webhook:path=/mutate-near-kotal-io-v1alpha1-node,mutating=true,failurePolicy=fail,groups=near.kotal.io,resources=nodes,verbs=create;update,versions=v1alpha1,name=mutate-near-v1alpha1-node.kb.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Defaulter = &Node{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (n *Node) Default() {
	nodelog.Info("default", "name", n.Name)

	if n.Spec.CPU == "" {
		n.Spec.CPU = DefaultNodeCPURequest
	}
	if n.Spec.CPULimit == "" {
		n.Spec.CPULimit = DefaultNodeCPULimit
	}

	if n.Spec.Memory == "" {
		n.Spec.Memory = DefaultNodeMemoryRequest
	}
	if n.Spec.MemoryLimit == "" {
		n.Spec.MemoryLimit = DefaultNodeMemoryLimit
	}

	if n.Spec.Storage == "" {
		// TODO: update with archival node defaulting
		n.Spec.Storage = DefaultNodeStorageRequest
	}

}