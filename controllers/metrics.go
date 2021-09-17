package controllers

import (
	ctx "context"
	"fmt"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type MetricsController struct{}

func (c MetricsController) NodeMetrics(context echo.Context) error {
	var client utils.Client = *utils.NewClient()
	metrics, err := client.MetricsV1beta1.NodeMetricses().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	nodeRowMetrics := RowNodeMetrics(metrics.Items, nodes.Items)
	return context.JSON(http.StatusOK, nodeRowMetrics)
}

func (c MetricsController) ClusterMetrics(context echo.Context) error {
	var client utils.Client = *utils.NewClient()
	metrics, err := client.MetricsV1beta1.NodeMetricses().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	ClusterRowMetrics := RowClusterMetrics(metrics.Items, nodes.Items)
	return context.JSON(http.StatusOK, ClusterRowMetrics)
}

func RowNodeMetrics(metrics []v1beta1.NodeMetrics, nodes []v1.Node) (rows []models.NodeRowMetrics) {
	for k, v := range nodes {
		var row models.NodeRowMetrics
		row.Name = v.ObjectMeta.Name
		row.Architecture = v.Status.NodeInfo.Architecture
		row.ContainerRuntimeVersion = v.Status.NodeInfo.ContainerRuntimeVersion
		row.KubeletVersion = v.Status.NodeInfo.KubeProxyVersion
		row.OperatingSystem = fmt.Sprintf("%s / %s", v.Status.NodeInfo.OperatingSystem, v.Status.NodeInfo.OSImage)
		row.Pods = v.Status.Allocatable.Pods().String()
		if len(v.Status.Addresses) > 1 {
			row.IpAddress = v.Status.Addresses[0].Address
			row.HostName = v.Status.Addresses[1].Address
		}

		row.TotalCpuCores = fmt.Sprintf("%vm", v.Status.Allocatable.Cpu().MilliValue())
		row.CpuUsageCores = fmt.Sprintf("%vm", metrics[k].Usage.Cpu().MilliValue())
		row.CpuUsagePercentage = fmt.Sprintf("%v%%", metrics[k].Usage.Cpu().MilliValue()*100/v.Status.Allocatable.Cpu().MilliValue())
		row.TotalMemory = fmt.Sprintf("%vMi", v.Status.Allocatable.Memory().Value()/(1024*1024))
		row.MemoryUsage = fmt.Sprintf("%vMi", metrics[k].Usage.Memory().Value()/(1024*1024))
		row.MemoryUsagePercentage = fmt.Sprintf("%v%%", (metrics[k].Usage.Memory().MilliValue()/(1024*1024))*100/(v.Status.Allocatable.Memory().MilliValue()/(1024*1024)))
		rows = append(rows, row)
	}

	return
}

func RowClusterMetrics(metrics []v1beta1.NodeMetrics, nodes []v1.Node) (rows models.ClusterMetricsCache) {
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

		row.TotalCpuCores = fmt.Sprintf("%vm", totalCpuCores)
		row.TotalCpuUsage = fmt.Sprintf("%vm", totalCpuUsage)
		row.TotalMemory = fmt.Sprintf("%vMi", totalMemory)
		row.TotalMemoryUsage = fmt.Sprintf("%vMi", totalMemoryUsage)
		row.Provider = nodes[0].Status.NodeInfo.KubeProxyVersion
		row.NodesCount = totalNodes
		row.CpuPercentage = fmt.Sprintf("%v%%", totalCpuUsage*100/totalCpuCores)
		row.MemoryPercentage = fmt.Sprintf("%v%%", totalMemoryUsage*100/totalMemory)
	}

	return row
}
