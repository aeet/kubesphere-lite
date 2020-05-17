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
	"my-kubesphere/pkg/servicemesh"
)

type Config struct {
	KubernetesOptions  *k8s.KubernetesOptions
	ServiceMeshOptions *servicemesh.Options
}

func New() *Config {
	return &Config{
		KubernetesOptions:  k8s.NewKubernetesOptions(),
		ServiceMeshOptions: servicemesh.NewServiceMeshOptions(),
	}
}
