package informers

import (
	snapshotclient "github.com/kubernetes-csi/external-snapshotter/v2/pkg/client/clientset/versioned"
	snapshotinformer "github.com/kubernetes-csi/external-snapshotter/v2/pkg/client/informers/externalversions"
	applicationclient "github.com/kubernetes-sigs/application/pkg/client/clientset/versioned"
	applicationinformers "github.com/kubernetes-sigs/application/pkg/client/informers/externalversions"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	istioinformers "istio.io/client-go/pkg/informers/externalversions"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apiextensionsinformers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	k8sinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"time"
)

const defaultResync = 600 * time.Second

type InformerFactory interface {
	KubernetesSharedInformerFactory() k8sinformers.SharedInformerFactory
	IstioSharedInformerFactory() istioinformers.SharedInformerFactory
	ApplicationSharedInformerFactory() applicationinformers.SharedInformerFactory
	SnapshotSharedInformerFactory() snapshotinformer.SharedInformerFactory
	ApiExtensionSharedInformerFactory() apiextensionsinformers.SharedInformerFactory
	Start(stopCh <-chan struct{})
}

type informerFactories struct {
	informerFactory              k8sinformers.SharedInformerFactory
	istioInformerFactory         istioinformers.SharedInformerFactory
	appInformerFactory           applicationinformers.SharedInformerFactory
	snapshotInformerFactory      snapshotinformer.SharedInformerFactory
	apiextensionsInformerFactory apiextensionsinformers.SharedInformerFactory
}

func NewInformerFactories(client kubernetes.Interface, istioClient istioclient.Interface,
	appClient applicationclient.Interface, snapshotClient snapshotclient.Interface, apiextensionsClient apiextensionsclient.Interface) InformerFactory {
	factory := &informerFactories{}

	if client != nil {
		factory.informerFactory = k8sinformers.NewSharedInformerFactory(client, defaultResync)
	}

	if appClient != nil {
		factory.appInformerFactory = applicationinformers.NewSharedInformerFactory(appClient, defaultResync)
	}

	if istioClient != nil {
		factory.istioInformerFactory = istioinformers.NewSharedInformerFactory(istioClient, defaultResync)
	}

	if snapshotClient != nil {
		factory.snapshotInformerFactory = snapshotinformer.NewSharedInformerFactory(snapshotClient, defaultResync)
	}

	if apiextensionsClient != nil {
		factory.apiextensionsInformerFactory = apiextensionsinformers.NewSharedInformerFactory(apiextensionsClient, defaultResync)
	}

	return factory
}

func (f *informerFactories) KubernetesSharedInformerFactory() k8sinformers.SharedInformerFactory {
	return f.informerFactory
}

func (f *informerFactories) ApplicationSharedInformerFactory() applicationinformers.SharedInformerFactory {
	return f.appInformerFactory
}

func (f *informerFactories) IstioSharedInformerFactory() istioinformers.SharedInformerFactory {
	return f.istioInformerFactory
}

func (f *informerFactories) SnapshotSharedInformerFactory() snapshotinformer.SharedInformerFactory {
	return f.snapshotInformerFactory
}

func (f *informerFactories) ApiExtensionSharedInformerFactory() apiextensionsinformers.SharedInformerFactory {
	return f.apiextensionsInformerFactory
}

func (f *informerFactories) Start(stopCh <-chan struct{}) {
	if f.informerFactory != nil {
		f.informerFactory.Start(stopCh)
	}

	if f.istioInformerFactory != nil {
		f.istioInformerFactory.Start(stopCh)
	}

	if f.appInformerFactory != nil {
		f.appInformerFactory.Start(stopCh)
	}

	if f.snapshotInformerFactory != nil {
		f.snapshotInformerFactory.Start(stopCh)
	}

	if f.apiextensionsInformerFactory != nil {
		f.apiextensionsInformerFactory.Start(stopCh)
	}
}
