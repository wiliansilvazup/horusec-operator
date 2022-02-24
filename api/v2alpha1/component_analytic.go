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
	"strconv"
)

func (h *HorusecPlatform) GetAnalyticComponent() Analytic {
	return h.Spec.Components.Analytic
}

func (h *HorusecPlatform) GetAnalyticAutoscaling() Autoscaling {
	return h.GetAnalyticComponent().Pod.Autoscaling
}

func (h *HorusecPlatform) GetAnalyticName() string {
	name := h.GetAnalyticComponent().Name
	if name == "" {
		return fmt.Sprintf("%s-analytic", h.GetName())
	}
	return name
}

func (h *HorusecPlatform) GetAnalyticPath() string {
	path := h.GetAnalyticComponent().Ingress.Path
	if path == "" {
		return "/analytic"
	}
	return path
}

func (h *HorusecPlatform) GetAnalyticPortHTTP() int {
	port := h.GetAnalyticComponent().Port.HTTP
	if port == 0 {
		return 8005
	}
	return port
}

func (h *HorusecPlatform) GetAnalyticLabels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       h.GetName(),
		"app.kubernetes.io/component":  "analytic",
		"app.kubernetes.io/managed-by": "horusec",
	}
}

func (h *HorusecPlatform) GetAnalyticV1ToV2Labels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       h.GetName(),
		"app.kubernetes.io/component":  "analytic-v1-2-v2",
		"app.kubernetes.io/managed-by": "horusec",
	}
}

func (h *HorusecPlatform) GetAnalyticReplicaCount() *int32 {
	if !h.GetAnalyticAutoscaling().Enabled {
		count := h.GetAnalyticComponent().ReplicaCount
		return &count
	}
	return nil
}

func (h *HorusecPlatform) GetAnalyticDefaultURL() string {
	return fmt.Sprintf("http://%s:%v", h.GetAnalyticName(), h.GetAnalyticPortHTTP())
}

func (h *HorusecPlatform) GetAnalyticRegistry() string {
	registry := h.GetAnalyticComponent().Container.Image.Registry
	if registry == "" {
		return "docker.io/"
	}
	return registry
}

func (h *HorusecPlatform) GetAnalyticRepository() string {
	repository := h.GetAnalyticComponent().Container.Image.Repository
	if repository == "" {
		return "wiliansilvazup/horusec-analytic"
	}
	return repository
}

func (h *HorusecPlatform) GetAnalyticTag() string {
	tag := h.GetAnalyticComponent().Container.Image.Tag
	if tag == "" {
		return h.GetLatestVersion()
	}
	return tag
}

func (h *HorusecPlatform) GetAnalyticImage() string {
	return fmt.Sprintf("%s/%s:%s", h.GetAnalyticRegistry(), h.GetAnalyticRepository(), h.GetAnalyticTag())
}

func (h *HorusecPlatform) GetAnalyticDatabaseLogMode() string {
	if h.Spec.Components.Analytic.Database.LogMode {
		return "true"
	}

	return "false"
}

func (h *HorusecPlatform) GetAnalyticDatabaseHost() string {
	host := h.Spec.Components.Analytic.Database.Host
	if host == "" {
		return "postgresql"
	}

	return host
}

func (h *HorusecPlatform) GetAnalyticDatabasePort() string {
	port := h.Spec.Components.Analytic.Database.Port
	if port <= 0 {
		return "5432"
	}

	return strconv.Itoa(port)
}

func (h *HorusecPlatform) GetAnalyticDatabaseName() string {
	name := h.Spec.Components.Analytic.Database.Name
	if name == "" {
		return "horusec_analytic_db"
	}

	return name
}

func (h *HorusecPlatform) GetAnalyticSSLMode() string {
	mode := h.Spec.Components.Analytic.Database.SslMode
	if mode == nil || *mode {
		return ""
	}

	return "?sslmode=disable"
}

func (h *HorusecPlatform) GetAnalyticDatabaseURI() string {
	return fmt.Sprintf("postgresql://$(HORUSEC_ANALYTIC_DATABASE_USERNAME):$(HORUSEC_ANALYTIC_DATABASE_PASSWORD)@%s:%s/%s%s",
		h.GetAnalyticDatabaseHost(), h.GetAnalyticDatabasePort(), h.GetAnalyticDatabaseName(), h.GetAnalyticSSLMode())
}

func (h *HorusecPlatform) GetAnalyticHost() string {
	host := h.Spec.Components.Analytic.Ingress.Host
	if host == "" {
		return "analytic.local"
	}

	return host
}

func (h *HorusecPlatform) IsAnalyticIngressEnabled() bool {
	enabled := h.Spec.Components.Analytic.Ingress.Enabled
	if enabled == nil {
		return true
	}

	return *enabled
}
