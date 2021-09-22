package controllers

import (
	ctx "context"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	v1 "k8s.io/api/apps/v1"
	core1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkLoadController struct{}

func (c WorkLoadController) Get(context echo.Context, nameSpaceName string) error {
	var client utils.Client = *utils.NewClient()
	deployments, deploymentErr := client.Clientset.AppsV1().Deployments(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{})
	pods, podsError := client.Clientset.CoreV1().Pods(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{})
	if deploymentErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: deploymentErr.Error(),
		})
	}
	if podsError != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: podsError.Error(),
		})
	}
	workLoads := getWorkLoad(deployments, pods)
	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(workLoads),
		ResourceType: utils.WORK_LOAD,
	})

}

func (c WorkLoadController) GetBySelector(context echo.Context, nameSpaceName string, selector string) error {
	var client utils.Client = *utils.NewClient()
	deployments, deploymentErr := client.Clientset.AppsV1().Deployments(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{
		LabelSelector: selector,
	})
	pods, podsError := client.Clientset.CoreV1().Pods(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{
		LabelSelector: selector,
	})
	if deploymentErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: deploymentErr.Error(),
		})
	}
	if podsError != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: podsError.Error(),
		})
	}
	workLoads := getWorkLoad(deployments, pods)
	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(workLoads),
		ResourceType: utils.WORK_LOAD,
	})
}

func getWorkLoad(deployments *v1.DeploymentList, pods *core1.PodList) []models.WorkLoad {
	var workLoads []models.WorkLoad
	var selectors []string
	for _, v := range deployments.Items {
		selectors = append(selectors, v.Labels["workload.user.cattle.io/workloadselector"])
		workLoads = append(workLoads, models.WorkLoad{
			Deployment: v,
			LinkedPods: getLinkedPods(pods, v.Labels["workload.user.cattle.io/workloadselector"]),
		})
	}

	otherPods := getUnLinked(pods, selectors)
	for _, v := range otherPods {
		workLoads = append(workLoads, models.WorkLoad{
			Pod: v,
		})
	}
	return workLoads
}

func getLinkedPods(pods *core1.PodList, wrokLoadselector string) []core1.Pod {
	var results []core1.Pod
	for _, v := range pods.Items {
		if v.Labels["workload.user.cattle.io/workloadselector"] == wrokLoadselector {
			results = append(results, v)
		}

	}
	return results
}

func getUnLinked(pods *core1.PodList, wrokLoadselectors []string) []core1.Pod {
	var results []core1.Pod

	for _, v := range pods.Items {
		for _, s := range wrokLoadselectors {
			if v.Labels["workload.user.cattle.io/workloadselector"] != s || v.Labels["workload.user.cattle.io/workloadselector"] == "" {
				results = append(results, v)
			}
		}
	}

	return results
}
