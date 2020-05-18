/*
 * @File: apis.go
 * @Date: 2020/5/18 11:47 上午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package apis

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/labels"
	"my-kubesphere/pkg/apiserver"
)

func GenerateHandlers(server apiserver.APIServer) *gin.Engine {
	r := gin.New()
	r.GET("/auth", func(context *gin.Context) {
		list, _ := server.InformerFactory.KubernetesSharedInformerFactory().Core().V1().Namespaces().Lister().List(labels.Everything())
		//list2, _ := server.InformerFactory.KubernetesSharedInformerFactory().Core().V1().Pods().Lister().List(labels.Everything())
		context.JSON(200, list)
	})
	return r
}
