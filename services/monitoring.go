package services

import (
	"encoding/json"

	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/sirupsen/logrus"
)

type MonitoringService struct {
	NameSpace string
	Resource  string
	EventName string
	PayLoad   interface{}
	ClusterId string
}

func (m MonitoringService) PushEvent(session *utils.Session) {
	m.ClusterId = utils.NewConfig().ClientId
	msg, _ := json.Marshal(m)
	err := session.Send(msg)
	if err != nil {
		logrus.Error(err)
	}

}
