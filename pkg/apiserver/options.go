/*
 * @File: options
 * @Date: 2020/5/17 8:02 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package apiserver

import (
	"github.com/emicklei/go-restful"
	"my-kubesphere/pkg/config"
	"my-kubesphere/pkg/informers"
	"my-kubesphere/pkg/k8s"
	"net/http"
)

type APIServer struct {
	ServerCount      int
	Server           *http.Server
	Config           *config.Config
	container        *restful.Container
	KubernetesClient k8s.Client
	InformerFactory  informers.InformerFactory
}
