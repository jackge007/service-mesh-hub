package kubeconfig

import (
	"context"

	"github.com/rotisserie/eris"
	k8s_core "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	FailedToFindKubeConfigSecret = func(err error, clusterName string) error {
		return eris.Wrapf(err, "Failed to find kube config secret for cluster %s", clusterName)
	}
	ClusterNotFound = func(clusterName string) error {
		return eris.Errorf("Cluster '%s' was not found", clusterName)
	}
)

func NewKubeConfigLookup(
	kubeClusterClient smh_discovery.KubernetesClusterClient,
	secretClient k8s_core.SecretClient,
	kubeConverter Converter,
) KubeConfigLookup {
	return &kubeConfigLookup{
		kubeClusterClient: kubeClusterClient,
		secretsClient:     secretClient,
		kubeConverter:     kubeConverter,
	}
}

type kubeConfigLookup struct {
	secretsClient     k8s_core.SecretClient
	kubeClusterClient smh_discovery.KubernetesClusterClient
	kubeConverter     Converter
}

func (k *kubeConfigLookup) FromCluster(ctx context.Context, clusterName string) (config *ConvertedConfigs, err error) {
	var kubeCluster *smh_discovery.KubernetesCluster
	allClusters, err := k.kubeClusterClient.ListKubernetesCluster(ctx)
	if err != nil {
		return nil, err
	}
	for _, foundCluster := range allClusters.Items {
		if foundCluster.GetName() == clusterName {
			kubeCluster = &foundCluster
			break
		}
	}

	if kubeCluster == nil {
		return nil, ClusterNotFound(clusterName)
	}

	cfgSecretRef := kubeCluster.Spec.GetSecretRef()
	secret, err := k.secretsClient.GetSecret(ctx, client.ObjectKey{Name: cfgSecretRef.GetName(), Namespace: cfgSecretRef.GetNamespace()})
	if err != nil {
		return nil, FailedToFindKubeConfigSecret(err, kubeCluster.GetName())
	}

	clusterName, config, err = k.kubeConverter.SecretToConfig(secret)
	if err != nil {
		return nil, err
	}

	return config, nil
}
