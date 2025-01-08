// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceMetadataServiceConfigurationInitParameters struct {

	// Indicates the minimum IMDS version that the notebook instance supports. When passed "1" is passed. This means that both IMDSv1 and IMDSv2 are supported. Valid values are 1 and 2.
	MinimumInstanceMetadataServiceVersion *string `json:"minimumInstanceMetadataServiceVersion,omitempty" tf:"minimum_instance_metadata_service_version,omitempty"`
}

type InstanceMetadataServiceConfigurationObservation struct {

	// Indicates the minimum IMDS version that the notebook instance supports. When passed "1" is passed. This means that both IMDSv1 and IMDSv2 are supported. Valid values are 1 and 2.
	MinimumInstanceMetadataServiceVersion *string `json:"minimumInstanceMetadataServiceVersion,omitempty" tf:"minimum_instance_metadata_service_version,omitempty"`
}

type InstanceMetadataServiceConfigurationParameters struct {

	// Indicates the minimum IMDS version that the notebook instance supports. When passed "1" is passed. This means that both IMDSv1 and IMDSv2 are supported. Valid values are 1 and 2.
	// +kubebuilder:validation:Optional
	MinimumInstanceMetadataServiceVersion *string `json:"minimumInstanceMetadataServiceVersion,omitempty" tf:"minimum_instance_metadata_service_version,omitempty"`
}

type NotebookInstanceInitParameters struct {

	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See Elastic Inference Accelerator for more details. Valid values: ml.eia1.medium, ml.eia1.large, ml.eia1.xlarge, ml.eia2.medium, ml.eia2.large, ml.eia2.xlarge.
	// +listType=set
	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty" tf:"accelerator_types,omitempty"`

	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in AWS CodeCommit or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	// +listType=set
	AdditionalCodeRepositories []*string `json:"additionalCodeRepositories,omitempty" tf:"additional_code_repositories,omitempty"`

	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in AWS CodeCommit or in any other Git repository.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta2.CodeRepository
	DefaultCodeRepository *string `json:"defaultCodeRepository,omitempty" tf:"default_code_repository,omitempty"`

	// Reference to a CodeRepository in sagemaker to populate defaultCodeRepository.
	// +kubebuilder:validation:Optional
	DefaultCodeRepositoryRef *v1.Reference `json:"defaultCodeRepositoryRef,omitempty" tf:"-"`

	// Selector for a CodeRepository in sagemaker to populate defaultCodeRepository.
	// +kubebuilder:validation:Optional
	DefaultCodeRepositorySelector *v1.Selector `json:"defaultCodeRepositorySelector,omitempty" tf:"-"`

	// Set to Disabled to disable internet access to notebook. Requires security_groups and subnet_id to be set. Supported values: Enabled (Default) or Disabled. If set to Disabled, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess *string `json:"directInternetAccess,omitempty" tf:"direct_internet_access,omitempty"`

	// Information on the IMDS configuration of the notebook instance. Conflicts with instance_metadata_service_configuration. see details below.
	InstanceMetadataServiceConfiguration *InstanceMetadataServiceConfigurationInitParameters `json:"instanceMetadataServiceConfiguration,omitempty" tf:"instance_metadata_service_configuration,omitempty"`

	// The name of ML compute instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName *string `json:"lifecycleConfigName,omitempty" tf:"lifecycle_config_name,omitempty"`

	// The platform identifier of the notebook instance runtime environment. This value can be either notebook-al1-v1, notebook-al2-v1, notebook-al2-v2, or notebook-al2-v3, depending on which version of Amazon Linux you require.
	PlatformIdentifier *string `json:"platformIdentifier,omitempty" tf:"platform_identifier,omitempty"`

	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Whether root access is Enabled or Disabled for users of the notebook instance. The default value is Enabled.
	RootAccess *string `json:"rootAccess,omitempty" tf:"root_access,omitempty"`

	// The associated security groups.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The VPC subnet ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

type NotebookInstanceObservation struct {

	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See Elastic Inference Accelerator for more details. Valid values: ml.eia1.medium, ml.eia1.large, ml.eia1.xlarge, ml.eia2.medium, ml.eia2.large, ml.eia2.xlarge.
	// +listType=set
	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty" tf:"accelerator_types,omitempty"`

	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in AWS CodeCommit or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	// +listType=set
	AdditionalCodeRepositories []*string `json:"additionalCodeRepositories,omitempty" tf:"additional_code_repositories,omitempty"`

	// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in AWS CodeCommit or in any other Git repository.
	DefaultCodeRepository *string `json:"defaultCodeRepository,omitempty" tf:"default_code_repository,omitempty"`

	// Set to Disabled to disable internet access to notebook. Requires security_groups and subnet_id to be set. Supported values: Enabled (Default) or Disabled. If set to Disabled, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess *string `json:"directInternetAccess,omitempty" tf:"direct_internet_access,omitempty"`

	// The name of the notebook instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Information on the IMDS configuration of the notebook instance. Conflicts with instance_metadata_service_configuration. see details below.
	InstanceMetadataServiceConfiguration *InstanceMetadataServiceConfigurationObservation `json:"instanceMetadataServiceConfiguration,omitempty" tf:"instance_metadata_service_configuration,omitempty"`

	// The name of ML compute instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName *string `json:"lifecycleConfigName,omitempty" tf:"lifecycle_config_name,omitempty"`

	// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting subnet_id.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The platform identifier of the notebook instance runtime environment. This value can be either notebook-al1-v1, notebook-al2-v1, notebook-al2-v2, or notebook-al2-v3, depending on which version of Amazon Linux you require.
	PlatformIdentifier *string `json:"platformIdentifier,omitempty" tf:"platform_identifier,omitempty"`

	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Whether root access is Enabled or Disabled for users of the notebook instance. The default value is Enabled.
	RootAccess *string `json:"rootAccess,omitempty" tf:"root_access,omitempty"`

	// The associated security groups.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The VPC subnet ID.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

type NotebookInstanceParameters struct {

	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See Elastic Inference Accelerator for more details. Valid values: ml.eia1.medium, ml.eia1.large, ml.eia1.xlarge, ml.eia2.medium, ml.eia2.large, ml.eia2.xlarge.
	// +kubebuilder:validation:Optional
	// +listType=set
	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty" tf:"accelerator_types,omitempty"`

	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in AWS CodeCommit or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	// +kubebuilder:validation:Optional
	// +listType=set
	AdditionalCodeRepositories []*string `json:"additionalCodeRepositories,omitempty" tf:"additional_code_repositories,omitempty"`

	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in AWS CodeCommit or in any other Git repository.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta2.CodeRepository
	// +kubebuilder:validation:Optional
	DefaultCodeRepository *string `json:"defaultCodeRepository,omitempty" tf:"default_code_repository,omitempty"`

	// Reference to a CodeRepository in sagemaker to populate defaultCodeRepository.
	// +kubebuilder:validation:Optional
	DefaultCodeRepositoryRef *v1.Reference `json:"defaultCodeRepositoryRef,omitempty" tf:"-"`

	// Selector for a CodeRepository in sagemaker to populate defaultCodeRepository.
	// +kubebuilder:validation:Optional
	DefaultCodeRepositorySelector *v1.Selector `json:"defaultCodeRepositorySelector,omitempty" tf:"-"`

	// Set to Disabled to disable internet access to notebook. Requires security_groups and subnet_id to be set. Supported values: Enabled (Default) or Disabled. If set to Disabled, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	// +kubebuilder:validation:Optional
	DirectInternetAccess *string `json:"directInternetAccess,omitempty" tf:"direct_internet_access,omitempty"`

	// Information on the IMDS configuration of the notebook instance. Conflicts with instance_metadata_service_configuration. see details below.
	// +kubebuilder:validation:Optional
	InstanceMetadataServiceConfiguration *InstanceMetadataServiceConfigurationParameters `json:"instanceMetadataServiceConfiguration,omitempty" tf:"instance_metadata_service_configuration,omitempty"`

	// The name of ML compute instance type.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The name of a lifecycle configuration to associate with the notebook instance.
	// +kubebuilder:validation:Optional
	LifecycleConfigName *string `json:"lifecycleConfigName,omitempty" tf:"lifecycle_config_name,omitempty"`

	// The platform identifier of the notebook instance runtime environment. This value can be either notebook-al1-v1, notebook-al2-v1, notebook-al2-v2, or notebook-al2-v3, depending on which version of Amazon Linux you require.
	// +kubebuilder:validation:Optional
	PlatformIdentifier *string `json:"platformIdentifier,omitempty" tf:"platform_identifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Whether root access is Enabled or Disabled for users of the notebook instance. The default value is Enabled.
	// +kubebuilder:validation:Optional
	RootAccess *string `json:"rootAccess,omitempty" tf:"root_access,omitempty"`

	// The associated security groups.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The VPC subnet ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

// NotebookInstanceSpec defines the desired state of NotebookInstance
type NotebookInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotebookInstanceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NotebookInstanceInitParameters `json:"initProvider,omitempty"`
}

// NotebookInstanceStatus defines the observed state of NotebookInstance.
type NotebookInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotebookInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NotebookInstance is the Schema for the NotebookInstances API. Provides a SageMaker Notebook Instance resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NotebookInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceType) || (has(self.initProvider) && has(self.initProvider.instanceType))",message="spec.forProvider.instanceType is a required parameter"
	Spec   NotebookInstanceSpec   `json:"spec"`
	Status NotebookInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotebookInstanceList contains a list of NotebookInstances
type NotebookInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebookInstance `json:"items"`
}

// Repository type metadata.
var (
	NotebookInstance_Kind             = "NotebookInstance"
	NotebookInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotebookInstance_Kind}.String()
	NotebookInstance_KindAPIVersion   = NotebookInstance_Kind + "." + CRDGroupVersion.String()
	NotebookInstance_GroupVersionKind = CRDGroupVersion.WithKind(NotebookInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&NotebookInstance{}, &NotebookInstanceList{})
}
