// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2alpha1

import (
	"fmt"
)

func (h *HorusecPlatform) GetCoreComponent() ExposableComponent {
	return h.Spec.Components.Core
}

func (h *HorusecPlatform) GetCoreAutoscaling() Autoscaling {
	return h.GetCoreComponent().Pod.Autoscaling
}

func (h *HorusecPlatform) GetCoreName() string {
	name := h.GetCoreComponent().Name
	if name == "" {
		return fmt.Sprintf("%s-core", h.GetName())
	}
	return name
}

func (h *HorusecPlatform) GetCorePath() string {
	path := h.GetCoreComponent().Ingress.Path
	if path == "" {
		return "/core"
	}
	return path
}

func (h *HorusecPlatform) GetCorePortHTTP() int {
	port := h.GetCoreComponent().Port.HTTP
	if port == 0 {
		return 8003
	}
	return port
}

func (h *HorusecPlatform) GetCoreLabels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       h.GetName(),
		"app.kubernetes.io/component":  "core",
		"app.kubernetes.io/managed-by": "horusec",
	}
}

func (h *HorusecPlatform) GetCoreReplicaCount() *int32 {
	if !h.GetCoreAutoscaling().Enabled {
		count := h.GetCoreComponent().ReplicaCount
		return &count
	}
	return nil
}

func (h *HorusecPlatform) GetCoreDefaultURL() string {
	return fmt.Sprintf("http://%s:%v", h.GetCoreName(), h.GetCorePortHTTP())
}

func (h *HorusecPlatform) GetCoreRegistry() string {
	registry := h.GetCoreComponent().Container.Image.Registry
	if registry == "" {
		return "docker.io/"
	}
	return registry
}

func (h *HorusecPlatform) GetCoreRepository() string {
	repository := h.GetCoreComponent().Container.Image.Repository
	if repository == "" {
		return "wiliansilvazup/horusec-core"
	}
	return repository
}

func (h *HorusecPlatform) GetCoreTag() string {
	tag := h.GetCoreComponent().Container.Image.Tag
	if tag == "" {
		return h.GetLatestVersion()
	}
	return tag
}

func (h *HorusecPlatform) GetCoreImage() string {
	return fmt.Sprintf("%s/%s:%s", h.GetCoreRegistry(), h.GetCoreRepository(), h.GetCoreTag())
}

func (h *HorusecPlatform) GetCoreHost() string {
	host := h.Spec.Components.Core.Ingress.Host
	if host == "" {
		return "core.local"
	}

	return host
}

func (h *HorusecPlatform) IsCoreIngressEnabled() bool {
	enabled := h.Spec.Components.Core.Ingress.Enabled
	if enabled == nil {
		return true
	}

	return *enabled
}
