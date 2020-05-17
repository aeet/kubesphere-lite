/*
 * @File: options
 * @Date: 2020/5/17 7:00 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */

package servicemesh

type Options struct {
	IstioPilotHost            string
	JaegerQueryHost           string
	ServicemeshPrometheusHost string
}

func NewServiceMeshOptions() *Options {
	return &Options{
		IstioPilotHost:            "",
		JaegerQueryHost:           "",
		ServicemeshPrometheusHost: "",
	}
}
