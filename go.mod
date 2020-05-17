module my-kubesphere

require (
	github.com/emicklei/go-restful v2.9.5+incompatible
	github.com/kiali/kiali v0.15.1-0.20191210080139-edbbad1ef779
	github.com/kubernetes-csi/external-snapshotter/v2 v2.1.1
	github.com/kubernetes-sigs/application v0.0.0-20191210100950-18cc93526ab4
	gopkg.in/yaml.v2 v2.2.8
	istio.io/client-go v0.0.0-20191113122552-9bd0ba57c3d2
	k8s.io/apiextensions-apiserver v0.17.3
	k8s.io/client-go v0.17.3

)

replace (
	github.com/emicklei/go-restful => github.com/emicklei/go-restful v2.9.5+incompatible
	github.com/kiali/kiali => github.com/kubesphere/kiali v0.15.1-0.20191210080139-edbbad1ef779
	github.com/kubernetes-csi/external-snapshotter/v2 => github.com/kubernetes-csi/external-snapshotter/v2 v2.1.0
	github.com/kubernetes-sigs/application => github.com/kubesphere/application v0.0.0-20191210100950-18cc93526ab4
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.4
	istio.io/client-go => istio.io/client-go v0.0.0-20191113122552-9bd0ba57c3d2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191114105449-027877536833
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191114101535-6c5935290e33
)

go 1.13
