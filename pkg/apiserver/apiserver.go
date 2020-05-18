/*
 * @File: apiserver
 * @Date: 2020/5/17 10:01 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package apiserver

import (
	"context"
	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog"
	"my-kubesphere/pkg/config"
	"my-kubesphere/pkg/informers"
	"my-kubesphere/pkg/k8s"
	"net/http"
)

var API *APIServer

type APIServer struct {
	ServerCount      int
	Server           *http.Server
	Config           *config.Config
	container        *restful.Container
	KubernetesClient k8s.Client
	InformerFactory  informers.InformerFactory
}

func (s *APIServer) Run(stopCh <-chan struct{}) (err error) {
	err = s.waitForResourceSync(stopCh)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-stopCh
		_ = s.Server.Shutdown(ctx)
	}()
	klog.V(0).Infof("Start listening on %s", s.Server.Addr)
	if s.Server.TLSConfig != nil {
		err = s.Server.ListenAndServeTLS("", "")
	} else {
		err = s.Server.ListenAndServe()
	}
	Api = s
	return err
}

func (s *APIServer) waitForResourceSync(stopCh <-chan struct{}) error {
	klog.V(0).Info("Start cache objects")
	discoveryClient := s.KubernetesClient.Kubernetes().Discovery()
	_, apiResourcesList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		return err
	}
	isResourceExists := func(resource schema.GroupVersionResource) bool {
		for _, apiResource := range apiResourcesList {
			if apiResource.GroupVersion == resource.GroupVersion().String() {
				for _, rsc := range apiResource.APIResources {
					if rsc.Name == resource.Resource {
						return true
					}
				}
			}
		}
		return false
	}
	k8sGVRs := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "namespaces"},
		{Group: "", Version: "v1", Resource: "nodes"},
		{Group: "", Version: "v1", Resource: "resourcequotas"},
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "persistentvolumeclaims"},
		{Group: "", Version: "v1", Resource: "secrets"},
		{Group: "", Version: "v1", Resource: "configmaps"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "roles"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "rolebindings"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "clusterroles"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "clusterrolebindings"},
		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "apps", Version: "v1", Resource: "daemonsets"},
		{Group: "apps", Version: "v1", Resource: "replicasets"},
		{Group: "apps", Version: "v1", Resource: "statefulsets"},
		{Group: "apps", Version: "v1", Resource: "controllerrevisions"},
		{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"},
		{Group: "batch", Version: "v1", Resource: "jobs"},
		{Group: "batch", Version: "v1beta1", Resource: "cronjobs"},
		{Group: "extensions", Version: "v1beta1", Resource: "ingresses"},
		{Group: "autoscaling", Version: "v2beta2", Resource: "horizontalpodautoscalers"},
		{Group: "networking.k8s.io", Version: "v1", Resource: "networkpolicies"},
	}
	for _, gvr := range k8sGVRs {
		if !isResourceExists(gvr) {
			klog.Warningf("resource %s not exists in the cluster", gvr)
		} else {
			_, err := s.InformerFactory.KubernetesSharedInformerFactory().ForResource(gvr)
			if err != nil {
				klog.Errorf("cannot create informer for %s", gvr)
				return err
			}
		}
	}
	// sharedInformer
	s.InformerFactory.KubernetesSharedInformerFactory().Start(stopCh)
	s.InformerFactory.KubernetesSharedInformerFactory().WaitForCacheSync(stopCh)
	// appInformer
	appInformerFactory := s.InformerFactory.ApplicationSharedInformerFactory()
	appGVRs := []schema.GroupVersionResource{
		{Group: "app.k8s.io", Version: "v1beta1", Resource: "applications"},
	}
	for _, gvr := range appGVRs {
		if !isResourceExists(gvr) {
			klog.Warningf("resource %s not exists in the cluster", gvr)
		} else {
			_, err = appInformerFactory.ForResource(gvr)
			if err != nil {
				return err
			}
		}
	}
	appInformerFactory.Start(stopCh)
	appInformerFactory.WaitForCacheSync(stopCh)
	// snapshotInformer
	snapshotInformerFactory := s.InformerFactory.SnapshotSharedInformerFactory()
	snapshotGVRs := []schema.GroupVersionResource{
		{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshotclasses"},
		{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshots"},
		{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshotcontents"},
	}
	for _, gvr := range snapshotGVRs {
		if !isResourceExists(gvr) {
			klog.Warningf("resource %s not exists in the cluster", gvr)
		} else {
			_, err = snapshotInformerFactory.ForResource(gvr)
			if err != nil {
				return err
			}
		}
	}
	snapshotInformerFactory.Start(stopCh)
	snapshotInformerFactory.WaitForCacheSync(stopCh)
	// apiextensionsInformer
	apiextensionsInformerFactory := s.InformerFactory.ApiExtensionSharedInformerFactory()
	apiextensionsGVRs := []schema.GroupVersionResource{
		{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"},
	}
	for _, gvr := range apiextensionsGVRs {
		if !isResourceExists(gvr) {
			klog.Warningf("resource %s not exists in the cluster", gvr)
		} else {
			_, err = apiextensionsInformerFactory.ForResource(gvr)
			if err != nil {
				return err
			}
		}
	}
	apiextensionsInformerFactory.Start(stopCh)
	apiextensionsInformerFactory.WaitForCacheSync(stopCh)
	klog.V(0).Info("Finished caching objects")
	return nil
}
