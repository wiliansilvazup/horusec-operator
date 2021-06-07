package api

import (
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
)

//nolint:lll, funlen // to improve in the future
func NewDeployment(resource *v2alpha1.HorusecPlatform) appsv1.Deployment {
	probe := corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/api/health",
				Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
			},
		},
	}
	return appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resource.GetAPIName(),
			Namespace: resource.GetNamespace(),
			Labels:    resource.GetApiLabels(),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: resource.GetAPIReplicaCount(),
			Selector: &metav1.LabelSelector{MatchLabels: resource.GetApiLabels()},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: resource.GetApiLabels()},
				Spec: corev1.PodSpec{Containers: []corev1.Container{{
					Name:  resource.GetAPIName(),
					Image: resource.GetAPIImage(),
					Env: []corev1.EnvVar{
						{Name: "HORUSEC_PORT", Value: strconv.Itoa(resource.GetAPIPortHTTP())},
						{Name: "HORUSEC_DATABASE_SQL_LOG_MODE", Value: resource.GetGlobalDatabaseLogMode()},
						{Name: "HORUSEC_GRPC_USE_CERTS", Value: strconv.FormatBool(resource.Spec.Global.GrpcUseCerts)},
						{Name: "HORUSEC_GRPC_AUTH_URL", Value: resource.GetAuthDefaultGRPCURL()},
						{Name: "HORUSEC_BROKER_HOST", Value: resource.GetGlobalBrokerHost()},
						{Name: "HORUSEC_BROKER_PORT", Value: resource.GetGlobalBrokerPort()},
						resource.NewEnvFromSecret("HORUSEC_BROKER_USERNAME", resource.GetGlobalBrokerUsername()),
						resource.NewEnvFromSecret("HORUSEC_BROKER_PASSWORD", resource.GetGlobalBrokerPassword()),
						resource.NewEnvFromSecret("HORUSEC_DATABASE_USERNAME", resource.GetGlobalDatabaseUsername()),
						resource.NewEnvFromSecret("HORUSEC_DATABASE_PASSWORD", resource.GetGlobalDatabasePassword()),
						{Name: "HORUSEC_DATABASE_SQL_URI", Value: resource.GetGlobalDatabaseURI()},
					},
					Ports: []corev1.ContainerPort{
						{Name: "http", ContainerPort: int32(resource.GetAPIPortHTTP())},
					},
					LivenessProbe:  &probe,
					ReadinessProbe: &probe,
				}}},
			},
		},
	}
}
