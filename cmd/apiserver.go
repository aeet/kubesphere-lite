package main

import (
	kiali "github.com/kiali/kiali/config"
	"my-kubesphere/pkg/config"
	"my-kubesphere/pkg/server"
)

type ServerRunOptions struct {
	ConfigFile          string
	Config              *config.Config
	GinServerRunOptions *server.GinServerRunOptions
	DebugMode           bool
}

func main() {
	s := &ServerRunOptions{
		ConfigFile:          "",
		GinServerRunOptions: server.NewServerRunOptions(),
		Config:              config.New(),
		DebugMode:           false,
	}
	kc := kiali.NewConfig()
	kc.API.Namespaces.Exclude = []string{"istio-system", "kubesphere*", "kube*"}
	kc.InCluster = true
	kc.ExternalServices.PrometheusServiceURL = s.Config.ServiceMeshOptions.ServicemeshPrometheusHost
	kc.ExternalServices.PrometheusCustomMetricsURL = kc.ExternalServices.PrometheusServiceURL
	kc.ExternalServices.Istio.UrlServiceVersion = s.Config.ServiceMeshOptions.IstioPilotHost
	kiali.Set(kc)
}
