package models

import (
	v1 "k8s.io/api/apps/v1"
	core1 "k8s.io/api/core/v1"
)

type WorkLoad struct {
	Deployment v1.Deployment `json:"deployment"`
	LinkedPods []core1.Pod   `json:"linkedPods"`
	Pod        core1.Pod     `json:"pod"`
}

type WorkLoadList struct {
	Items []WorkLoad `json:"items"`
}
