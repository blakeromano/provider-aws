/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package gatewayroute

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	tjcontroller "github.com/upbound/upjet/pkg/controller"
	"github.com/upbound/upjet/pkg/terraform"
	ctrl "sigs.k8s.io/controller-runtime"

	v1beta1 "github.com/upbound/provider-aws/apis/appmesh/v1beta1"
)

// Setup adds a controller that reconciles GatewayRoute managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1beta1.GatewayRoute_GroupVersionKind.String())
	var initializers managed.InitializerChain
	for _, i := range o.Provider.Resources["aws_appmesh_gateway_route"].InitializerFns {
		initializers = append(initializers, i(mgr.GetClient()))
	}
	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK))
	}
	r := managed.NewReconciler(mgr,
		xpresource.ManagedKind(v1beta1.GatewayRoute_GroupVersionKind),
		managed.WithExternalConnecter(tjcontroller.NewConnector(mgr.GetClient(), o.WorkspaceStore, o.SetupFn, o.Provider.Resources["aws_appmesh_gateway_route"],
			tjcontroller.WithCallbackProvider(tjcontroller.NewAPICallbacks(mgr, xpresource.ManagedKind(v1beta1.GatewayRoute_GroupVersionKind))),
		)),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(terraform.NewWorkspaceFinalizer(o.WorkspaceStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3*time.Minute),
		managed.WithInitializers(initializers),
		managed.WithConnectionPublishers(cps...),
		managed.WithPollInterval(o.PollInterval),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.GatewayRoute{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}
