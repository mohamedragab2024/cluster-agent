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
	fmt.Print("Getting metrics for nodes")
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
	nodeRowMetrics := ToRow(metrics.Items, nodes.Items)
	fmt.Print("metrics: ", nodeRowMetrics)
	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(nodeRowMetrics),
		ResourceType: utils.RESOUCETYPE_NODES,
	})
}

func ToRow(metrics []v1beta1.NodeMetrics, nodes []v1.Node) (rows []models.NodeRowMetrics) {
	for k, v := range nodes {
		var row models.NodeRowMetrics
		row.Architecture = v.Status.NodeInfo.Architecture
		row.ContainerRuntimeVersion = v.Status.NodeInfo.ContainerRuntimeVersion
		row.KubeletVersion = v.Status.NodeInfo.KubeProxyVersion
		row.OperatingSystem = fmt.Sprintf("%s / %s", v.Status.NodeInfo.OperatingSystem, v.Status.NodeInfo.OSImage)
		row.Pods = v.Status.Allocatable.Pods().String()
		if len(v.Status.Addresses) > 1 {
			row.IpAddress = v.Status.Addresses[0].Address
			row.HostName = v.Status.Addresses[1].Address
		}
		row.TotalCpuCors = fmt.Sprintf("%vm", v.Status.Allocatable.Cpu().MilliValue())
		row.TotalMemory = fmt.Sprintf("%vMi", v.Status.Allocatable.Memory().MilliValue())
		row.CpuUsageCors = fmt.Sprintf("%vm", metrics[k].Usage.Cpu().MilliValue())
		row.MemoryUsage = fmt.Sprintf("%vMi", metrics[k].Usage.Memory().MilliValue())
		rows = append(rows, row)
	}

	return
}
