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
)

func GenerateHandlers() *gin.Engine {
	r := gin.New()
	r.GET("/auth")
	return r
}
