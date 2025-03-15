/*
Copyright 2016 The Kubernetes Authors.

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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	_ "k8s.io/component-base/metrics/prometheus/restclient" // for client-go metrics registration

	k8smetrics "k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
)

const (
	caNamespace = "cluster_autoscaler"
)

var (
	podBlockingReason = k8smetrics.NewGaugeVec(
		&k8smetrics.GaugeOpts{
			Namespace: caNamespace,
			Name:      "pod_blocking_node",
			Help:      "Whether a pod is blocking a node from being removed",
		},
		[]string{"namespace", "pod", "reason"},
	)
)

// RegisterAll registers all metrics.
func RegisterAll(emitPodBlockingMetrics bool) {
	if emitPodBlockingMetrics {
		legacyregistry.MustRegister(podBlockingReason)
	}
}

// SetPodBlockingReason sets the pod blocking reason for the given pod
func SetPodBlockingReason(namespace, pod, reason string) {
	podBlockingReason.WithLabelValues(namespace, pod, string(reason)).Set(1)
}

// ClearPodBlockingReason clears the pod blocking reason for the given pod
func ClearPodBlockingReason(namespace, pod string) {
	podBlockingReason.DeletePartialMatch(prometheus.Labels{
		"namespace": namespace,
		"pod":       pod,
	})
}
