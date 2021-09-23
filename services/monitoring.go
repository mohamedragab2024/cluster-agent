package services

import (
	"encoding/json"

	utils "github.com/kube-carbonara/cluster-agent/utils"
)

type MonitoringService struct {
	NameSpace string
	Resource  string
	EventName string
	PayLoad   interface{}
	ClusterId string
}

func (m MonitoringService) PushEvent(session *utils.Session) error {
	m.ClusterId = utils.NewConfig().ClientId
	msg, _ := json.Marshal(m)
	err := session.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
