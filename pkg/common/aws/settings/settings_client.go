package settings

import (
	"context"

	smh_settings "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1"
	smh_settings_types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	container_runtime "github.com/solo-io/service-mesh-hub/pkg/common/container-runtime"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/metadata"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type awsSettingsHelperClient struct {
	settingsClient smh_settings.SettingsClient
}

func NewAwsSettingsHelperClient(settingsClient smh_settings.SettingsClient) SettingsHelperClient {
	return &awsSettingsHelperClient{settingsClient: settingsClient}
}

func (a *awsSettingsHelperClient) getSettingsSpec(ctx context.Context) (*smh_settings_types.SettingsSpec, error) {
	settings, err := a.settingsClient.GetSettings(
		ctx,
		client.ObjectKey{Name: metadata.GlobalSettingsName, Namespace: container_runtime.GetWriteNamespace()},
	)
	if errors.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &settings.Spec, nil
	}
}

func (a *awsSettingsHelperClient) GetAWSSettingsForAccount(
	ctx context.Context,
	accountId string,
) (*smh_settings_types.SettingsSpec_AwsAccount, error) {
	settingsSpec, err := a.getSettingsSpec(ctx)
	if err != nil {
		return nil, err
	}
	if settingsSpec.GetAws().GetDisabled() {
		return nil, nil
	}
	for _, awsAccountConfig := range settingsSpec.GetAws().GetAccounts() {
		if awsAccountConfig.GetAccountId() == accountId {
			return awsAccountConfig, nil
		}
	}
	return nil, nil
}
