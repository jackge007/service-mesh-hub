package mesh_networking

import (
	"context"

	"github.com/solo-io/smh/pkg/common/bootstrap"
	"github.com/solo-io/smh/pkg/mesh-networking/reporting"
	"github.com/solo-io/smh/pkg/mesh-networking/translation/istio"
	"github.com/solo-io/smh/pkg/mesh-networking/validation"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/snapshot/input"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/smh/pkg/mesh-networking/reconciliation"
)

// the mesh-networking controller is the Kubernetes Controller/Operator
// which processes k8s storage events to produce
// discovered resources.
func Start(ctx context.Context, opts bootstrap.Options) error {
	return bootstrap.Start(ctx, "networking", startReconciler, opts)
}

// start the main reconcile loop
func startReconciler(
	ctx context.Context,
	masterManager manager.Manager,
	mcClient multicluster.Client,
	_ multicluster.ClusterSet,
	_ multicluster.ClusterWatcher,
) error {

	snapshotBuilder := input.NewSingleClusterBuilder(masterManager.GetClient())
	reporter := reporting.NewPanickingReporter(ctx)
	istioTranslator := istio.NewIstioTranslator()
	validator := validation.NewValidator(istioTranslator)

	return reconciliation.Start(
		ctx,
		snapshotBuilder,
		validator,
		reporter,
		istioTranslator,
		mcClient,
		masterManager,
	)
}
