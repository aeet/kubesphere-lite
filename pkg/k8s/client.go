/*
 * @File: client
 * @Date: 2020/5/17 8:06 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package k8s

import (
	snapshotclient "github.com/kubernetes-csi/external-snapshotter/v2/pkg/client/clientset/versioned"
	applicationclientset "github.com/kubernetes-sigs/application/pkg/client/clientset/versioned"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client interface {
	Kubernetes() kubernetes.Interface
	Istio() istioclient.Interface
	Application() applicationclientset.Interface
	Snapshot() snapshotclient.Interface
	ApiExtensions() apiextensionsclient.Interface
	Discovery() discovery.DiscoveryInterface
	Master() string
	Config() *rest.Config
}

type kubernetesClient struct {
	k8s             kubernetes.Interface
	discoveryClient *discovery.DiscoveryClient
	application     applicationclientset.Interface
	istio           istioclient.Interface
	snapshot        snapshotclient.Interface
	apiextensions   apiextensionsclient.Interface
	master          string
	config          *rest.Config
}

func NewKubernetesClient(options *KubernetesOptions) (Client, error) {
	config, err := clientcmd.BuildConfigFromFlags("", options.KubeConfig)
	if err != nil {
		return nil, err
	}
	config.QPS = options.QPS
	config.Burst = options.Burst

	var k kubernetesClient
	k.k8s, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	k.application, err = applicationclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	k.istio, err = istioclient.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	k.snapshot, err = snapshotclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	k.apiextensions, err = apiextensionsclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	k.master = options.Master
	k.config = config

	return &k, nil
}

func (k *kubernetesClient) Kubernetes() kubernetes.Interface {
	return k.k8s
}

func (k *kubernetesClient) Discovery() discovery.DiscoveryInterface {
	return k.discoveryClient
}

func (k *kubernetesClient) Application() applicationclientset.Interface {
	return k.application
}

func (k *kubernetesClient) Istio() istioclient.Interface {
	return k.istio
}

func (k *kubernetesClient) Snapshot() snapshotclient.Interface {
	return k.snapshot
}

func (k *kubernetesClient) ApiExtensions() apiextensionsclient.Interface {
	return k.apiextensions
}

func (k *kubernetesClient) Master() string {
	return k.master
}

func (k *kubernetesClient) Config() *rest.Config {
	return k.config
}
