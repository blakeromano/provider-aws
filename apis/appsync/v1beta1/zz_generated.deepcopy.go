//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKey) DeepCopyInto(out *APIKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKey.
func (in *APIKey) DeepCopy() *APIKey {
	if in == nil {
		return nil
	}
	out := new(APIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyList) DeepCopyInto(out *APIKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyList.
func (in *APIKeyList) DeepCopy() *APIKeyList {
	if in == nil {
		return nil
	}
	out := new(APIKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyObservation) DeepCopyInto(out *APIKeyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyObservation.
func (in *APIKeyObservation) DeepCopy() *APIKeyObservation {
	if in == nil {
		return nil
	}
	out := new(APIKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyParameters) DeepCopyInto(out *APIKeyParameters) {
	*out = *in
	if in.APIID != nil {
		in, out := &in.APIID, &out.APIID
		*out = new(string)
		**out = **in
	}
	if in.APIIDRef != nil {
		in, out := &in.APIIDRef, &out.APIIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.APIIDSelector != nil {
		in, out := &in.APIIDSelector, &out.APIIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expires != nil {
		in, out := &in.Expires, &out.Expires
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyParameters.
func (in *APIKeyParameters) DeepCopy() *APIKeyParameters {
	if in == nil {
		return nil
	}
	out := new(APIKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeySpec) DeepCopyInto(out *APIKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeySpec.
func (in *APIKeySpec) DeepCopy() *APIKeySpec {
	if in == nil {
		return nil
	}
	out := new(APIKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyStatus) DeepCopyInto(out *APIKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyStatus.
func (in *APIKeyStatus) DeepCopy() *APIKeyStatus {
	if in == nil {
		return nil
	}
	out := new(APIKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalAuthenticationProviderObservation) DeepCopyInto(out *AdditionalAuthenticationProviderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalAuthenticationProviderObservation.
func (in *AdditionalAuthenticationProviderObservation) DeepCopy() *AdditionalAuthenticationProviderObservation {
	if in == nil {
		return nil
	}
	out := new(AdditionalAuthenticationProviderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalAuthenticationProviderParameters) DeepCopyInto(out *AdditionalAuthenticationProviderParameters) {
	*out = *in
	if in.AuthenticationType != nil {
		in, out := &in.AuthenticationType, &out.AuthenticationType
		*out = new(string)
		**out = **in
	}
	if in.LambdaAuthorizerConfig != nil {
		in, out := &in.LambdaAuthorizerConfig, &out.LambdaAuthorizerConfig
		*out = make([]LambdaAuthorizerConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OpenIDConnectConfig != nil {
		in, out := &in.OpenIDConnectConfig, &out.OpenIDConnectConfig
		*out = make([]OpenIDConnectConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserPoolConfig != nil {
		in, out := &in.UserPoolConfig, &out.UserPoolConfig
		*out = make([]UserPoolConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalAuthenticationProviderParameters.
func (in *AdditionalAuthenticationProviderParameters) DeepCopy() *AdditionalAuthenticationProviderParameters {
	if in == nil {
		return nil
	}
	out := new(AdditionalAuthenticationProviderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPI) DeepCopyInto(out *GraphQLAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPI.
func (in *GraphQLAPI) DeepCopy() *GraphQLAPI {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphQLAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPILambdaAuthorizerConfigObservation) DeepCopyInto(out *GraphQLAPILambdaAuthorizerConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPILambdaAuthorizerConfigObservation.
func (in *GraphQLAPILambdaAuthorizerConfigObservation) DeepCopy() *GraphQLAPILambdaAuthorizerConfigObservation {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPILambdaAuthorizerConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPILambdaAuthorizerConfigParameters) DeepCopyInto(out *GraphQLAPILambdaAuthorizerConfigParameters) {
	*out = *in
	if in.AuthorizerResultTTLInSeconds != nil {
		in, out := &in.AuthorizerResultTTLInSeconds, &out.AuthorizerResultTTLInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.AuthorizerURI != nil {
		in, out := &in.AuthorizerURI, &out.AuthorizerURI
		*out = new(string)
		**out = **in
	}
	if in.IdentityValidationExpression != nil {
		in, out := &in.IdentityValidationExpression, &out.IdentityValidationExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPILambdaAuthorizerConfigParameters.
func (in *GraphQLAPILambdaAuthorizerConfigParameters) DeepCopy() *GraphQLAPILambdaAuthorizerConfigParameters {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPILambdaAuthorizerConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIList) DeepCopyInto(out *GraphQLAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GraphQLAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIList.
func (in *GraphQLAPIList) DeepCopy() *GraphQLAPIList {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphQLAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIObservation) DeepCopyInto(out *GraphQLAPIObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIObservation.
func (in *GraphQLAPIObservation) DeepCopy() *GraphQLAPIObservation {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIOpenIDConnectConfigObservation) DeepCopyInto(out *GraphQLAPIOpenIDConnectConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIOpenIDConnectConfigObservation.
func (in *GraphQLAPIOpenIDConnectConfigObservation) DeepCopy() *GraphQLAPIOpenIDConnectConfigObservation {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIOpenIDConnectConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIOpenIDConnectConfigParameters) DeepCopyInto(out *GraphQLAPIOpenIDConnectConfigParameters) {
	*out = *in
	if in.AuthTTL != nil {
		in, out := &in.AuthTTL, &out.AuthTTL
		*out = new(float64)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.IatTTL != nil {
		in, out := &in.IatTTL, &out.IatTTL
		*out = new(float64)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIOpenIDConnectConfigParameters.
func (in *GraphQLAPIOpenIDConnectConfigParameters) DeepCopy() *GraphQLAPIOpenIDConnectConfigParameters {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIOpenIDConnectConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIParameters) DeepCopyInto(out *GraphQLAPIParameters) {
	*out = *in
	if in.AdditionalAuthenticationProvider != nil {
		in, out := &in.AdditionalAuthenticationProvider, &out.AdditionalAuthenticationProvider
		*out = make([]AdditionalAuthenticationProviderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthenticationType != nil {
		in, out := &in.AuthenticationType, &out.AuthenticationType
		*out = new(string)
		**out = **in
	}
	if in.LambdaAuthorizerConfig != nil {
		in, out := &in.LambdaAuthorizerConfig, &out.LambdaAuthorizerConfig
		*out = make([]GraphQLAPILambdaAuthorizerConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LogConfig != nil {
		in, out := &in.LogConfig, &out.LogConfig
		*out = make([]LogConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OpenIDConnectConfig != nil {
		in, out := &in.OpenIDConnectConfig, &out.OpenIDConnectConfig
		*out = make([]GraphQLAPIOpenIDConnectConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
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
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UserPoolConfig != nil {
		in, out := &in.UserPoolConfig, &out.UserPoolConfig
		*out = make([]GraphQLAPIUserPoolConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.XrayEnabled != nil {
		in, out := &in.XrayEnabled, &out.XrayEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIParameters.
func (in *GraphQLAPIParameters) DeepCopy() *GraphQLAPIParameters {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPISpec) DeepCopyInto(out *GraphQLAPISpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPISpec.
func (in *GraphQLAPISpec) DeepCopy() *GraphQLAPISpec {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIStatus) DeepCopyInto(out *GraphQLAPIStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIStatus.
func (in *GraphQLAPIStatus) DeepCopy() *GraphQLAPIStatus {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIUserPoolConfigObservation) DeepCopyInto(out *GraphQLAPIUserPoolConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIUserPoolConfigObservation.
func (in *GraphQLAPIUserPoolConfigObservation) DeepCopy() *GraphQLAPIUserPoolConfigObservation {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIUserPoolConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLAPIUserPoolConfigParameters) DeepCopyInto(out *GraphQLAPIUserPoolConfigParameters) {
	*out = *in
	if in.AppIDClientRegex != nil {
		in, out := &in.AppIDClientRegex, &out.AppIDClientRegex
		*out = new(string)
		**out = **in
	}
	if in.AwsRegion != nil {
		in, out := &in.AwsRegion, &out.AwsRegion
		*out = new(string)
		**out = **in
	}
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
	if in.UserPoolIDRef != nil {
		in, out := &in.UserPoolIDRef, &out.UserPoolIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.UserPoolIDSelector != nil {
		in, out := &in.UserPoolIDSelector, &out.UserPoolIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLAPIUserPoolConfigParameters.
func (in *GraphQLAPIUserPoolConfigParameters) DeepCopy() *GraphQLAPIUserPoolConfigParameters {
	if in == nil {
		return nil
	}
	out := new(GraphQLAPIUserPoolConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaAuthorizerConfigObservation) DeepCopyInto(out *LambdaAuthorizerConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaAuthorizerConfigObservation.
func (in *LambdaAuthorizerConfigObservation) DeepCopy() *LambdaAuthorizerConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LambdaAuthorizerConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaAuthorizerConfigParameters) DeepCopyInto(out *LambdaAuthorizerConfigParameters) {
	*out = *in
	if in.AuthorizerResultTTLInSeconds != nil {
		in, out := &in.AuthorizerResultTTLInSeconds, &out.AuthorizerResultTTLInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.AuthorizerURI != nil {
		in, out := &in.AuthorizerURI, &out.AuthorizerURI
		*out = new(string)
		**out = **in
	}
	if in.IdentityValidationExpression != nil {
		in, out := &in.IdentityValidationExpression, &out.IdentityValidationExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaAuthorizerConfigParameters.
func (in *LambdaAuthorizerConfigParameters) DeepCopy() *LambdaAuthorizerConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LambdaAuthorizerConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfigObservation) DeepCopyInto(out *LogConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfigObservation.
func (in *LogConfigObservation) DeepCopy() *LogConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LogConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfigParameters) DeepCopyInto(out *LogConfigParameters) {
	*out = *in
	if in.CloudwatchLogsRoleArn != nil {
		in, out := &in.CloudwatchLogsRoleArn, &out.CloudwatchLogsRoleArn
		*out = new(string)
		**out = **in
	}
	if in.CloudwatchLogsRoleArnRef != nil {
		in, out := &in.CloudwatchLogsRoleArnRef, &out.CloudwatchLogsRoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudwatchLogsRoleArnSelector != nil {
		in, out := &in.CloudwatchLogsRoleArnSelector, &out.CloudwatchLogsRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExcludeVerboseContent != nil {
		in, out := &in.ExcludeVerboseContent, &out.ExcludeVerboseContent
		*out = new(bool)
		**out = **in
	}
	if in.FieldLogLevel != nil {
		in, out := &in.FieldLogLevel, &out.FieldLogLevel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfigParameters.
func (in *LogConfigParameters) DeepCopy() *LogConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LogConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnectConfigObservation) DeepCopyInto(out *OpenIDConnectConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnectConfigObservation.
func (in *OpenIDConnectConfigObservation) DeepCopy() *OpenIDConnectConfigObservation {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnectConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnectConfigParameters) DeepCopyInto(out *OpenIDConnectConfigParameters) {
	*out = *in
	if in.AuthTTL != nil {
		in, out := &in.AuthTTL, &out.AuthTTL
		*out = new(float64)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.IatTTL != nil {
		in, out := &in.IatTTL, &out.IatTTL
		*out = new(float64)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnectConfigParameters.
func (in *OpenIDConnectConfigParameters) DeepCopy() *OpenIDConnectConfigParameters {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnectConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolConfigObservation) DeepCopyInto(out *UserPoolConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolConfigObservation.
func (in *UserPoolConfigObservation) DeepCopy() *UserPoolConfigObservation {
	if in == nil {
		return nil
	}
	out := new(UserPoolConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolConfigParameters) DeepCopyInto(out *UserPoolConfigParameters) {
	*out = *in
	if in.AppIDClientRegex != nil {
		in, out := &in.AppIDClientRegex, &out.AppIDClientRegex
		*out = new(string)
		**out = **in
	}
	if in.AwsRegion != nil {
		in, out := &in.AwsRegion, &out.AwsRegion
		*out = new(string)
		**out = **in
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolConfigParameters.
func (in *UserPoolConfigParameters) DeepCopy() *UserPoolConfigParameters {
	if in == nil {
		return nil
	}
	out := new(UserPoolConfigParameters)
	in.DeepCopyInto(out)
	return out
}
