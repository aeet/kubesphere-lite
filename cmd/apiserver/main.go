/*
 * @File: options
 * @Date: 2020/5/17 7:00 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */

package main

import (
	"crypto/tls"
	"fmt"
	kiali "github.com/kiali/kiali/config"
	"my-kubesphere/pkg/apis"
	"my-kubesphere/pkg/apiserver"
	"my-kubesphere/pkg/config"
	"my-kubesphere/pkg/informers"
	"my-kubesphere/pkg/k8s"
	"my-kubesphere/pkg/signals"
	"net/http"
)

type RunOptions struct {
	ConfigFile          string
	Config              *config.Config
	GinServerRunOptions *apis.GinServerRunOptions
	DebugMode           bool
}

func main() {
	s := &RunOptions{
		ConfigFile:          "",
		GinServerRunOptions: apis.NewGinServerRunOptions(),
		Config:              config.New(),
		DebugMode:           false,
	}
	initKiali(s)
	er := Run(s, signals.SetupSignalHandler())
	if er != nil {
		fmt.Println(er)
	}
}

func initKiali(s *RunOptions) {
	kc := kiali.NewConfig()
	kc.API.Namespaces.Exclude = []string{"istio-system", "kubesphere*", "kube*"}
	kc.InCluster = true
	kc.ExternalServices.PrometheusServiceURL = s.Config.ServiceMeshOptions.ServicemeshPrometheusHost
	kc.ExternalServices.PrometheusCustomMetricsURL = kc.ExternalServices.PrometheusServiceURL
	kc.ExternalServices.Istio.UrlServiceVersion = s.Config.ServiceMeshOptions.IstioPilotHost
	kiali.Set(kc)
}

func Run(s *RunOptions, stopCh <-chan struct{}) error {
	apiServer, err := s.NewAPIServer()
	if err != nil {
		return err
	}
	apiServer.InstallAPI()
	return apiServer.Run(stopCh)
}

func (s *RunOptions) NewAPIServer() (*apiserver.APIServer, error) {
	apiServer := &apiserver.APIServer{
		Config: s.Config,
	}

	kubernetesClient, err := k8s.NewKubernetesClient(s.Config.KubernetesOptions)
	if err != nil {
		return nil, err
	}
	apiServer.KubernetesClient = kubernetesClient

	informerFactory := informers.NewInformerFactories(kubernetesClient.Kubernetes(), kubernetesClient.Istio(), kubernetesClient.Application(), kubernetesClient.Snapshot(), kubernetesClient.ApiExtensions())
	apiServer.InformerFactory = informerFactory

	webServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.GinServerRunOptions.InsecurePort),
		Handler: apis.GenerateHandlers(),
	}

	if s.GinServerRunOptions.SecurePort != 0 {
		certificate, err := tls.LoadX509KeyPair(s.GinServerRunOptions.TlsCertFile, s.GinServerRunOptions.TlsPrivateKey)
		if err != nil {
			return nil, err
		}
		webServer.TLSConfig.Certificates = []tls.Certificate{certificate}
	}
	apiServer.Server = webServer
	return apiServer, nil
}
