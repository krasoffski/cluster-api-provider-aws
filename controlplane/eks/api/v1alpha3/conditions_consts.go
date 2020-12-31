/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

import clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"

const (
	// EKSControlPlaneReadyCondition condition reports on the successful reconciliation of eks control plane.
	EKSControlPlaneReadyCondition clusterv1.ConditionType = "EKSControlPlaneReady"
	// EKSControlPlaneReconciliationFailedReason used to report failures while reconciling EKS control plane
	EKSControlPlaneReconciliationFailedReason = "EKSControlPlaneReconciliationFailed"
)

const (
	// IAMControlPlaneRolesReadyCondition condition reports on the successful reconciliation of eks control plane iam roles.
	IAMControlPlaneRolesReadyCondition clusterv1.ConditionType = "IAMControlPlaneRolesReady"
	// IAMControlPlaneRolesReconciliationFailedReason used to report failures while reconciling EKS control plane iam roles
	IAMControlPlaneRolesReconciliationFailedReason = "IAMControlPlaneRolesReconciliationFailed"
)

const (
	// IAMAuthenticatorConfiguredCondition condition reports on the successful reconciliation of aws-iam-authenticator config
	IAMAuthenticatorConfiguredCondition clusterv1.ConditionType = "IAMAuthenticatorConfigured"
	// IAMAuthenticatorConfigurationFailedReason used to report failures while reconciling the aws-iam-authenticator config
	IAMAuthenticatorConfigurationFailedReason = "IAMAuthenticatorConfigurationFailed"
)

const (
	// EKSAddonsConfiguredCondition condition reports on the successful reconciliation of EKS addons
	EKSAddonsConfiguredCondition clusterv1.ConditionType = "EKSAddonsConfigured"
	// EKSAddonsConfiguredFailedReason used to report failures while reconciling the EKS addons
	EKSAddonsConfiguredFailedReason = "EKSAddonsConfiguredFailed"
)
