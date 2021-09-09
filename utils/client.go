package utils

import (
	"k8s.io/client-go/kubernetes"
	networkingv1client "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
)

type Client struct {
	Clientset          *kubernetes.Clientset
	Networkingv1client *networkingv1client.NetworkingV1Client
}

func NewClient() *Client {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, _ := kubernetes.NewForConfig(config)
	ntClient, _ := networkingv1client.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &Client{
		Clientset:          clientset,
		Networkingv1client: ntClient,
	}
}
