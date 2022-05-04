// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodegroupSpec defines the desired state of Nodegroup.
//
// An object representing an Amazon EKS managed node group.
type NodegroupSpec struct {
	// The AMI type for your node group. GPU instance types should use the AL2_x86_64_GPU
	// AMI type. Non-GPU instances should use the AL2_x86_64 AMI type. Arm instances
	// should use the AL2_ARM_64 AMI type. All types use the Amazon EKS optimized
	// Amazon Linux 2 AMI. If you specify launchTemplate, and your launch template
	// uses a custom AMI, then don't specify amiType, or the node group deployment
	// will fail. For more information about using launch templates with Amazon
	// EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	AMIType *string `json:"amiType,omitempty"`
	// The capacity type for your node group.
	CapacityType *string `json:"capacityType,omitempty"`
	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `json:"clientRequestToken,omitempty"`
	// The name of the cluster to create the node group in.
	ClusterName *string                                  `json:"clusterName,omitempty"`
	ClusterRef  *ackv1alpha1.AWSResourceReferenceWrapper `json:"clusterRef,omitempty"`
	// The root device disk size (in GiB) for your node group instances. The default
	// disk size is 20 GiB. If you specify launchTemplate, then don't specify diskSize,
	// or the node group deployment will fail. For more information about using
	// launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	DiskSize *int64 `json:"diskSize,omitempty"`
	// Specify the instance types for a node group. If you specify a GPU instance
	// type, be sure to specify AL2_x86_64_GPU with the amiType parameter. If you
	// specify launchTemplate, then you can specify zero or one instance type in
	// your launch template or you can specify 0-20 instance types for instanceTypes.
	// If however, you specify an instance type in your launch template and specify
	// any instanceTypes, the node group deployment will fail. If you don't specify
	// an instance type in a launch template or for instanceTypes, then t3.medium
	// is used, by default. If you specify Spot for capacityType, then we recommend
	// specifying multiple values for instanceTypes. For more information, see Managed
	// node group capacity types (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types)
	// and Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	InstanceTypes []*string `json:"instanceTypes,omitempty"`
	// The Kubernetes labels to be applied to the nodes in the node group when they
	// are created.
	Labels map[string]*string `json:"labels,omitempty"`
	// An object representing a node group's launch template specification. If specified,
	// then do not specify instanceTypes, diskSize, or remoteAccess and make sure
	// that the launch template meets the requirements in launchTemplateSpecification.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate,omitempty"`
	// The unique name to give your node group.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node
	// group. The Amazon EKS worker node kubelet daemon makes calls to Amazon Web
	// Services APIs on your behalf. Nodes receive permissions for these API calls
	// through an IAM instance profile and associated policies. Before you can launch
	// nodes and register them into a cluster, you must create an IAM role for those
	// nodes to use when they are launched. For more information, see Amazon EKS
	// node IAM role (https://docs.aws.amazon.com/eks/latest/userguide/worker_node_IAM_role.html)
	// in the Amazon EKS User Guide . If you specify launchTemplate, then don't
	// specify IamInstanceProfile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html)
	// in your launch template, or the node group deployment will fail. For more
	// information about using launch templates with Amazon EKS, see Launch template
	// support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	// +kubebuilder:validation:Required
	NodeRole *string `json:"nodeRole"`
	// The AMI version of the Amazon EKS optimized AMI to use with your node group.
	// By default, the latest available AMI version for the node group's current
	// Kubernetes version is used. For more information, see Amazon EKS optimized
	// Amazon Linux 2 AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
	// in the Amazon EKS User Guide. If you specify launchTemplate, and your launch
	// template uses a custom AMI, then don't specify releaseVersion, or the node
	// group deployment will fail. For more information about using launch templates
	// with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	ReleaseVersion *string `json:"releaseVersion,omitempty"`
	// The remote access (SSH) configuration to use with your node group. If you
	// specify launchTemplate, then don't specify remoteAccess, or the node group
	// deployment will fail. For more information about using launch templates with
	// Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	RemoteAccess *RemoteAccessConfig `json:"remoteAccess,omitempty"`
	// The scaling configuration details for the Auto Scaling group that is created
	// for your node group.
	ScalingConfig *NodegroupScalingConfig `json:"scalingConfig,omitempty"`
	// The subnets to use for the Auto Scaling group that is created for your node
	// group. If you specify launchTemplate, then don't specify SubnetId (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html)
	// in your launch template, or the node group deployment will fail. For more
	// information about using launch templates with Amazon EKS, see Launch template
	// support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	// +kubebuilder:validation:Required
	Subnets []*string `json:"subnets"`
	// The metadata to apply to the node group to assist with categorization and
	// organization. Each tag consists of a key and an optional value. You define
	// both. Node group tags do not propagate to any other resources associated
	// with the node group, such as the Amazon EC2 instances or subnets.
	Tags map[string]*string `json:"tags,omitempty"`
	// The Kubernetes taints to be applied to the nodes in the node group.
	Taints []*Taint `json:"taints,omitempty"`
	// The node group update configuration.
	UpdateConfig *NodegroupUpdateConfig `json:"updateConfig,omitempty"`
	// The Kubernetes version to use for your managed nodes. By default, the Kubernetes
	// version of the cluster is used, and this is the only accepted specified value.
	// If you specify launchTemplate, and your launch template uses a custom AMI,
	// then don't specify version, or the node group deployment will fail. For more
	// information about using launch templates with Amazon EKS, see Launch template
	// support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	Version *string `json:"version,omitempty"`
}

// NodegroupStatus defines the observed state of Nodegroup
type NodegroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The Unix epoch timestamp in seconds for when the managed node group was created.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// The health status of the node group. If there are issues with your node group's
	// health, they are listed here.
	// +kubebuilder:validation:Optional
	Health *NodegroupHealth `json:"health,omitempty"`
	// The Unix epoch timestamp in seconds for when the managed node group was last
	// modified.
	// +kubebuilder:validation:Optional
	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// The resources associated with the node group, such as Auto Scaling groups
	// and security groups for remote access.
	// +kubebuilder:validation:Optional
	Resources *NodegroupResources `json:"resources,omitempty"`
	// The current status of the managed node group.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// Nodegroup is the Schema for the Nodegroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Nodegroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodegroupSpec   `json:"spec,omitempty"`
	Status            NodegroupStatus `json:"status,omitempty"`
}

// NodegroupList contains a list of Nodegroup
// +kubebuilder:object:root=true
type NodegroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Nodegroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Nodegroup{}, &NodegroupList{})
}
