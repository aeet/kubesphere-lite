/*
 * @File: namespace
 * @Date: 2020/5/18 2:28 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package namespace

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/labels"
	"my-kubesphere/pkg/apiserver"
)

func NameSpaceList(ctx *gin.Context) {
	list, _ := apiserver.API.InformerFactory.KubernetesSharedInformerFactory().Core().V1().Namespaces().Lister().List(labels.Everything())
	ctx.JSON(200, list)
}
