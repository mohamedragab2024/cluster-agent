package handlers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
)

type WatcherHanlder struct{}

func (c WatcherHanlder) Handle() {
	(controllers.ServicesController{}).Watch()
	//go (controllers.PodsController{}).Watch()
}
