module my-kubesphere

require (
	github.com/emicklei/go-restful v2.12.0+incompatible
	github.com/kiali/kiali v0.15.1-0.20191210080139-edbbad1ef779
	github.com/kubernetes-csi/external-snapshotter/v2 v2.1.1
	github.com/kubernetes-sigs/application v0.0.0-20191210100950-18cc93526ab4
	gopkg.in/yaml.v2 v2.3.0 // indirect
	istio.io/client-go v0.0.0-20200513000250-b1d6e9886b7b
	k8s.io/apiextensions-apiserver v0.18.2
	k8s.io/client-go v11.0.0+incompatible

)

replace (
	github.com/kiali/kiali => github.com/kubesphere/kiali v0.15.1-0.20191210080139-edbbad1ef779
	github.com/kubernetes-sigs/application => github.com/kubesphere/application v0.0.0-20191210100950-18cc93526ab4

)

go 1.13
