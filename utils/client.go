package utils

import (
	"fmt"

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
		fmt.Print(err.Error())
		panic(err.Error())
	}
	clientset, _ := kubernetes.NewForConfig(config)
	ntClient, _ := networkingv1client.NewForConfig(config)
	if err != nil {
		fmt.Print(err.Error())
		panic(err.Error())
	}
	return &Client{
		Clientset:          clientset,
		Networkingv1client: ntClient,
	}
}
