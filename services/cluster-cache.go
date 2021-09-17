package services

import (
	"bytes"
	ctx "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type ClusterCacheService struct{}

func (c ClusterCacheService) PushMetricsUpdates() {
	metrics, err := c.ClusterMetrics()
	if err != nil {
		log.Println("write:", err)
		return
	}
	fmt.Print("Update cluster metrics cache ...", metrics)

	jsonReq, err := json.Marshal(metrics)
	if err != nil {
		log.Println("write:", err)
		return
	}

	config := utils.NewConfig()
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/clusters/%s", config.RemoteProxy, config.ClientId), bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Println("write:", err)
		return
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("x-agent", config.ClientId)
	r.Header.Add("x-agent-app-key", config.AppKey)
	resp, _ := client.Do(r)
	if err != nil {
		log.Println("write:", err)
		return
	}
	fmt.Print(resp)

}

func (c ClusterCacheService) ClusterMetrics() (models.ClusterMetricsCache, error) {
	var client utils.Client = *utils.NewClient()
	metrics, err := client.MetricsV1beta1.NodeMetricses().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("write:", err)
		return models.ClusterMetricsCache{}, err
	}
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("write:", err)
		return models.ClusterMetricsCache{}, err
	}
	ClusterRowMetrics := c.RowClusterMetrics(metrics.Items, nodes.Items)
	return ClusterRowMetrics, nil
}

func (c ClusterCacheService) RowClusterMetrics(metrics []v1beta1.NodeMetrics, nodes []v1.Node) (rows models.ClusterMetricsCache) {
	var row models.ClusterMetricsCache
	if len(nodes) > 0 {
		var totalCpuCores int64 = 0
		var totalCpuUsage int64 = 0
		var totalMemory int64 = 0
		var totalMemoryUsage int64 = 0
		var totalNodes int64 = 0
		for k, v := range nodes {
			totalCpuCores += v.Status.Allocatable.Cpu().MilliValue()
			totalCpuUsage += metrics[k].Usage.Cpu().MilliValue()
			totalMemory += v.Status.Allocatable.Memory().Value() / (1024 * 1024)
			totalMemoryUsage += metrics[k].Usage.Memory().Value() / (1024 * 1024)
			totalNodes++
		}

		row.TotalCpuCores = totalCpuCores
		row.TotalCpuUsage = totalCpuUsage
		row.TotalMemory = totalMemory
		row.TotalMemoryUsage = totalMemoryUsage
		row.Provider = nodes[0].Status.NodeInfo.KubeProxyVersion
		row.NodesCount = totalNodes
		row.CpuPercentage = fmt.Sprintf("%v%%", totalCpuUsage*100/totalCpuCores)
		row.MemoryPercentage = fmt.Sprintf("%v%%", totalMemoryUsage*100/totalMemory)
	}

	return row
}
