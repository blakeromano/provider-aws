/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Database.
func (mg *Database) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &v1beta1.BucketList{},
			Managed: &v1beta1.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamedQuery.
func (mg *NamedQuery) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Database),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseRef,
		Selector:     mg.Spec.ForProvider.DatabaseSelector,
		To: reference.To{
			List:    &DatabaseList{},
			Managed: &Database{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Database")
	}
	mg.Spec.ForProvider.Database = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Workgroup),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkgroupRef,
		Selector:     mg.Spec.ForProvider.WorkgroupSelector,
		To: reference.To{
			List:    &WorkgroupList{},
			Managed: &Workgroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Workgroup")
	}
	mg.Spec.ForProvider.Workgroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkgroupRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workgroup.
func (mg *Workgroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Configuration[i3].ResultConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnRef,
					Selector:     mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnSelector,
					To: reference.To{
						List:    &v1beta11.KeyList{},
						Managed: &v1beta11.Key{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn")
				}
				mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}