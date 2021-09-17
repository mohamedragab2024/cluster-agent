package utils

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	networkingv1client "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
	metricsv1alpha1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1alpha1"
	metricsv1beta1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

type Client struct {
	Clientset          *kubernetes.Clientset
	Networkingv1client *networkingv1client.NetworkingV1Client
	MetricsV1alpha1    *metricsv1alpha1.MetricsV1alpha1Client
	MetricsV1beta1     *metricsv1beta1.MetricsV1beta1Client
}

func NewClient() *Client {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Print(err.Error())
		panic(err.Error())
	}
	clientset, _ := kubernetes.NewForConfig(config)
	ntClient, _ := networkingv1client.NewForConfig(config)
	mtClientBeta, _ := metricsv1beta1.NewForConfig(config)
	mtClientAlpha, _ := metricsv1alpha1.NewForConfig(config)
	if err != nil {
		fmt.Print(err.Error())
		panic(err.Error())
	}
	return &Client{
		Clientset:          clientset,
		Networkingv1client: ntClient,
		MetricsV1alpha1:    mtClientAlpha,
		MetricsV1beta1:     mtClientBeta,
	}
}
