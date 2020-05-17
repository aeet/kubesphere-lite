/*
 * @File: options
 * @Date: 2020/5/17 7:00 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */

package k8s

type KubernetesOptions struct {
	KubeConfig string
	Master     string
	QPS        float32
	Burst      int
}

func NewKubernetesOptions() *KubernetesOptions {
	return &KubernetesOptions{
		KubeConfig: "",
		QPS:        1e6,
		Burst:      1e6,
	}
}
