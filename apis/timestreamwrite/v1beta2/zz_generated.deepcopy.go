//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositePartitionKeyInitParameters) DeepCopyInto(out *CompositePartitionKeyInitParameters) {
	*out = *in
	if in.EnforcementInRecord != nil {
		in, out := &in.EnforcementInRecord, &out.EnforcementInRecord
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositePartitionKeyInitParameters.
func (in *CompositePartitionKeyInitParameters) DeepCopy() *CompositePartitionKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(CompositePartitionKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositePartitionKeyObservation) DeepCopyInto(out *CompositePartitionKeyObservation) {
	*out = *in
	if in.EnforcementInRecord != nil {
		in, out := &in.EnforcementInRecord, &out.EnforcementInRecord
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositePartitionKeyObservation.
func (in *CompositePartitionKeyObservation) DeepCopy() *CompositePartitionKeyObservation {
	if in == nil {
		return nil
	}
	out := new(CompositePartitionKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositePartitionKeyParameters) DeepCopyInto(out *CompositePartitionKeyParameters) {
	*out = *in
	if in.EnforcementInRecord != nil {
		in, out := &in.EnforcementInRecord, &out.EnforcementInRecord
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositePartitionKeyParameters.
func (in *CompositePartitionKeyParameters) DeepCopy() *CompositePartitionKeyParameters {
	if in == nil {
		return nil
	}
	out := new(CompositePartitionKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MagneticStoreRejectedDataLocationInitParameters) DeepCopyInto(out *MagneticStoreRejectedDataLocationInitParameters) {
	*out = *in
	if in.S3Configuration != nil {
		in, out := &in.S3Configuration, &out.S3Configuration
		*out = new(S3ConfigurationInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MagneticStoreRejectedDataLocationInitParameters.
func (in *MagneticStoreRejectedDataLocationInitParameters) DeepCopy() *MagneticStoreRejectedDataLocationInitParameters {
	if in == nil {
		return nil
	}
	out := new(MagneticStoreRejectedDataLocationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MagneticStoreRejectedDataLocationObservation) DeepCopyInto(out *MagneticStoreRejectedDataLocationObservation) {
	*out = *in
	if in.S3Configuration != nil {
		in, out := &in.S3Configuration, &out.S3Configuration
		*out = new(S3ConfigurationObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MagneticStoreRejectedDataLocationObservation.
func (in *MagneticStoreRejectedDataLocationObservation) DeepCopy() *MagneticStoreRejectedDataLocationObservation {
	if in == nil {
		return nil
	}
	out := new(MagneticStoreRejectedDataLocationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MagneticStoreRejectedDataLocationParameters) DeepCopyInto(out *MagneticStoreRejectedDataLocationParameters) {
	*out = *in
	if in.S3Configuration != nil {
		in, out := &in.S3Configuration, &out.S3Configuration
		*out = new(S3ConfigurationParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MagneticStoreRejectedDataLocationParameters.
func (in *MagneticStoreRejectedDataLocationParameters) DeepCopy() *MagneticStoreRejectedDataLocationParameters {
	if in == nil {
		return nil
	}
	out := new(MagneticStoreRejectedDataLocationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MagneticStoreWritePropertiesInitParameters) DeepCopyInto(out *MagneticStoreWritePropertiesInitParameters) {
	*out = *in
	if in.EnableMagneticStoreWrites != nil {
		in, out := &in.EnableMagneticStoreWrites, &out.EnableMagneticStoreWrites
		*out = new(bool)
		**out = **in
	}
	if in.MagneticStoreRejectedDataLocation != nil {
		in, out := &in.MagneticStoreRejectedDataLocation, &out.MagneticStoreRejectedDataLocation
		*out = new(MagneticStoreRejectedDataLocationInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MagneticStoreWritePropertiesInitParameters.
func (in *MagneticStoreWritePropertiesInitParameters) DeepCopy() *MagneticStoreWritePropertiesInitParameters {
	if in == nil {
		return nil
	}
	out := new(MagneticStoreWritePropertiesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MagneticStoreWritePropertiesObservation) DeepCopyInto(out *MagneticStoreWritePropertiesObservation) {
	*out = *in
	if in.EnableMagneticStoreWrites != nil {
		in, out := &in.EnableMagneticStoreWrites, &out.EnableMagneticStoreWrites
		*out = new(bool)
		**out = **in
	}
	if in.MagneticStoreRejectedDataLocation != nil {
		in, out := &in.MagneticStoreRejectedDataLocation, &out.MagneticStoreRejectedDataLocation
		*out = new(MagneticStoreRejectedDataLocationObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MagneticStoreWritePropertiesObservation.
func (in *MagneticStoreWritePropertiesObservation) DeepCopy() *MagneticStoreWritePropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(MagneticStoreWritePropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MagneticStoreWritePropertiesParameters) DeepCopyInto(out *MagneticStoreWritePropertiesParameters) {
	*out = *in
	if in.EnableMagneticStoreWrites != nil {
		in, out := &in.EnableMagneticStoreWrites, &out.EnableMagneticStoreWrites
		*out = new(bool)
		**out = **in
	}
	if in.MagneticStoreRejectedDataLocation != nil {
		in, out := &in.MagneticStoreRejectedDataLocation, &out.MagneticStoreRejectedDataLocation
		*out = new(MagneticStoreRejectedDataLocationParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MagneticStoreWritePropertiesParameters.
func (in *MagneticStoreWritePropertiesParameters) DeepCopy() *MagneticStoreWritePropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(MagneticStoreWritePropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPropertiesInitParameters) DeepCopyInto(out *RetentionPropertiesInitParameters) {
	*out = *in
	if in.MagneticStoreRetentionPeriodInDays != nil {
		in, out := &in.MagneticStoreRetentionPeriodInDays, &out.MagneticStoreRetentionPeriodInDays
		*out = new(float64)
		**out = **in
	}
	if in.MemoryStoreRetentionPeriodInHours != nil {
		in, out := &in.MemoryStoreRetentionPeriodInHours, &out.MemoryStoreRetentionPeriodInHours
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPropertiesInitParameters.
func (in *RetentionPropertiesInitParameters) DeepCopy() *RetentionPropertiesInitParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionPropertiesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPropertiesObservation) DeepCopyInto(out *RetentionPropertiesObservation) {
	*out = *in
	if in.MagneticStoreRetentionPeriodInDays != nil {
		in, out := &in.MagneticStoreRetentionPeriodInDays, &out.MagneticStoreRetentionPeriodInDays
		*out = new(float64)
		**out = **in
	}
	if in.MemoryStoreRetentionPeriodInHours != nil {
		in, out := &in.MemoryStoreRetentionPeriodInHours, &out.MemoryStoreRetentionPeriodInHours
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPropertiesObservation.
func (in *RetentionPropertiesObservation) DeepCopy() *RetentionPropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionPropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPropertiesParameters) DeepCopyInto(out *RetentionPropertiesParameters) {
	*out = *in
	if in.MagneticStoreRetentionPeriodInDays != nil {
		in, out := &in.MagneticStoreRetentionPeriodInDays, &out.MagneticStoreRetentionPeriodInDays
		*out = new(float64)
		**out = **in
	}
	if in.MemoryStoreRetentionPeriodInHours != nil {
		in, out := &in.MemoryStoreRetentionPeriodInHours, &out.MemoryStoreRetentionPeriodInHours
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPropertiesParameters.
func (in *RetentionPropertiesParameters) DeepCopy() *RetentionPropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionPropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigurationInitParameters) DeepCopyInto(out *S3ConfigurationInitParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.EncryptionOption != nil {
		in, out := &in.EncryptionOption, &out.EncryptionOption
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.ObjectKeyPrefix != nil {
		in, out := &in.ObjectKeyPrefix, &out.ObjectKeyPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigurationInitParameters.
func (in *S3ConfigurationInitParameters) DeepCopy() *S3ConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(S3ConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigurationObservation) DeepCopyInto(out *S3ConfigurationObservation) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.EncryptionOption != nil {
		in, out := &in.EncryptionOption, &out.EncryptionOption
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.ObjectKeyPrefix != nil {
		in, out := &in.ObjectKeyPrefix, &out.ObjectKeyPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigurationObservation.
func (in *S3ConfigurationObservation) DeepCopy() *S3ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(S3ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigurationParameters) DeepCopyInto(out *S3ConfigurationParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.EncryptionOption != nil {
		in, out := &in.EncryptionOption, &out.EncryptionOption
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.ObjectKeyPrefix != nil {
		in, out := &in.ObjectKeyPrefix, &out.ObjectKeyPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigurationParameters.
func (in *S3ConfigurationParameters) DeepCopy() *S3ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(S3ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaInitParameters) DeepCopyInto(out *SchemaInitParameters) {
	*out = *in
	if in.CompositePartitionKey != nil {
		in, out := &in.CompositePartitionKey, &out.CompositePartitionKey
		*out = new(CompositePartitionKeyInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaInitParameters.
func (in *SchemaInitParameters) DeepCopy() *SchemaInitParameters {
	if in == nil {
		return nil
	}
	out := new(SchemaInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaObservation) DeepCopyInto(out *SchemaObservation) {
	*out = *in
	if in.CompositePartitionKey != nil {
		in, out := &in.CompositePartitionKey, &out.CompositePartitionKey
		*out = new(CompositePartitionKeyObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaObservation.
func (in *SchemaObservation) DeepCopy() *SchemaObservation {
	if in == nil {
		return nil
	}
	out := new(SchemaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaParameters) DeepCopyInto(out *SchemaParameters) {
	*out = *in
	if in.CompositePartitionKey != nil {
		in, out := &in.CompositePartitionKey, &out.CompositePartitionKey
		*out = new(CompositePartitionKeyParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaParameters.
func (in *SchemaParameters) DeepCopy() *SchemaParameters {
	if in == nil {
		return nil
	}
	out := new(SchemaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableInitParameters) DeepCopyInto(out *TableInitParameters) {
	*out = *in
	if in.MagneticStoreWriteProperties != nil {
		in, out := &in.MagneticStoreWriteProperties, &out.MagneticStoreWriteProperties
		*out = new(MagneticStoreWritePropertiesInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionProperties != nil {
		in, out := &in.RetentionProperties, &out.RetentionProperties
		*out = new(RetentionPropertiesInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(SchemaInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableInitParameters.
func (in *TableInitParameters) DeepCopy() *TableInitParameters {
	if in == nil {
		return nil
	}
	out := new(TableInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableList) DeepCopyInto(out *TableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableList.
func (in *TableList) DeepCopy() *TableList {
	if in == nil {
		return nil
	}
	out := new(TableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableObservation) DeepCopyInto(out *TableObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MagneticStoreWriteProperties != nil {
		in, out := &in.MagneticStoreWriteProperties, &out.MagneticStoreWriteProperties
		*out = new(MagneticStoreWritePropertiesObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionProperties != nil {
		in, out := &in.RetentionProperties, &out.RetentionProperties
		*out = new(RetentionPropertiesObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(SchemaObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableObservation.
func (in *TableObservation) DeepCopy() *TableObservation {
	if in == nil {
		return nil
	}
	out := new(TableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableParameters) DeepCopyInto(out *TableParameters) {
	*out = *in
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.DatabaseNameRef != nil {
		in, out := &in.DatabaseNameRef, &out.DatabaseNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseNameSelector != nil {
		in, out := &in.DatabaseNameSelector, &out.DatabaseNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MagneticStoreWriteProperties != nil {
		in, out := &in.MagneticStoreWriteProperties, &out.MagneticStoreWriteProperties
		*out = new(MagneticStoreWritePropertiesParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RetentionProperties != nil {
		in, out := &in.RetentionProperties, &out.RetentionProperties
		*out = new(RetentionPropertiesParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(SchemaParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableParameters.
func (in *TableParameters) DeepCopy() *TableParameters {
	if in == nil {
		return nil
	}
	out := new(TableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpec) DeepCopyInto(out *TableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpec.
func (in *TableSpec) DeepCopy() *TableSpec {
	if in == nil {
		return nil
	}
	out := new(TableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableStatus) DeepCopyInto(out *TableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableStatus.
func (in *TableStatus) DeepCopy() *TableStatus {
	if in == nil {
		return nil
	}
	out := new(TableStatus)
	in.DeepCopyInto(out)
	return out
}