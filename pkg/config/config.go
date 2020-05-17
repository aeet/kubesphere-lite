/*
 * @File: options
 * @Date: 2020/5/17 7:00 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package config

import (
	"my-kubesphere/pkg/k8s"
	"my-kubesphere/pkg/server"
)

type Config struct {
	ServerRunOptions  *server.ServerRunOptions
	KubernetesOptions *k8s.KubernetesOptions
}

func New() *Config {
	return &Config{
		ServerRunOptions:  server.NewServerRunOptions(),
		KubernetesOptions: k8s.NewKubernetesOptions(),
	}
}
