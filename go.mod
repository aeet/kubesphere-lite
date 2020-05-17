module my-kubesphere

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/kiali/kiali v0.15.1-0.20191210080139-edbbad1ef779
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace github.com/kiali/kiali => github.com/kubesphere/kiali v0.15.1-0.20191210080139-edbbad1ef779

go 1.13
