module github.com/kube-carbonara/cluster-agent

go 1.16

require (
	github.com/gorilla/websocket v1.4.2
	github.com/labstack/echo/v4 v4.5.0
	github.com/rancher/remotedialer v0.2.5
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	k8s.io/metrics v0.22.1
)
