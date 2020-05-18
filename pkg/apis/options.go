/*
 * @File: options
 * @Date: 2020/5/17 7:00 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */

package apis

type GinServerRunOptions struct {
	BindAddress   string
	InsecurePort  int
	SecurePort    int
	TlsCertFile   string
	TlsPrivateKey string
}

func NewGinServerRunOptions() *GinServerRunOptions {
	s := GinServerRunOptions{
		BindAddress:   "0.0.0.0",
		InsecurePort:  9090,
		SecurePort:    0,
		TlsCertFile:   "",
		TlsPrivateKey: "",
	}
	return &s
}
