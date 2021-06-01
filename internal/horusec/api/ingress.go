package api

import (
	"k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
)

func NewIngressRule(resource *v2alpha1.HorusecPlatform, pathType v1beta1.PathType) v1beta1.IngressRule {
	if !resource.Spec.Components.Api.Ingress.Enabled {
		return v1beta1.IngressRule{}
	}

	return v1beta1.IngressRule{
		Host: resource.Spec.Components.Api.Ingress.Host,
		IngressRuleValue: v1beta1.IngressRuleValue{
			HTTP: &v1beta1.HTTPIngressRuleValue{
				Paths: []v1beta1.HTTPIngressPath{
					{
						Path:     resource.Spec.Components.Api.Ingress.Path,
						PathType: &pathType,
						Backend: v1beta1.IngressBackend{
							ServiceName: resource.Spec.Components.Api.Name,
							ServicePort: intstr.IntOrString{
								Type:   0,
								IntVal: int32(resource.Spec.Components.Api.Port.HTTP),
							},
						},
					},
				},
			},
		},
	}
}

func NewIngressTLS(resource *v2alpha1.HorusecPlatform) v1beta1.IngressTLS {
	if !resource.Spec.Components.Api.Ingress.Enabled {
		return v1beta1.IngressTLS{}
	}

	return v1beta1.IngressTLS{
		Hosts:      []string{resource.Spec.Components.Api.Ingress.Host},
		SecretName: resource.Spec.Components.Api.Ingress.TLS.SecretName,
	}
}
