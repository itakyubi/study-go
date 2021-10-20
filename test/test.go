package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("https://127.0.0.1:6443", "/Users/wuao/.kube/config")
	clientset, _ := kubernetes.NewForConfig(config)

	informerFactory := informers.NewSharedInformerFactoryWithOptions(clientset, time.Minute, informers.WithNamespace("default"))
	informer := informerFactory.Core().V1().ConfigMaps()
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})
	lister := informer.Lister()

	stopCh := make(chan struct{})
	defer close(stopCh)
	informerFactory.Start(stopCh)

	if !cache.WaitForCacheSync(stopCh, informer.Informer().HasSynced) {
		return
	}

	configMaps, _ := lister.ConfigMaps("default").List(labels.Everything())
	for _, configMap := range configMaps {
		fmt.Printf("%s\r\n", configMap.Name)
	}
	<-stopCh
}

func onAdd(obj interface{}) {
	/*configMap := obj.(*v1.ConfigMap)
	fmt.Printf("onAdd:%s\r\n", configMap.Name)
	configMapBytes, _ := json.Marshal(configMap)
	fmt.Printf("onAdd:%s\r\n", string(configMapBytes))*/
}

func onUpdate(old, new interface{}) {
	oldConfigMap := old.(*v1.ConfigMap)
	newConfigMap := new.(*v1.ConfigMap)
	if oldConfigMap.Name != "helmet-detection-1-rtsp-reader1" {
		return
	}

	//fmt.Printf("onUpdate:%s to %s\r\n", oldConfigMap.Name, newConfigMap.Name)
	oldConfigMapBytes, _ := json.Marshal(oldConfigMap)
	newConfigMapBytes, _ := json.Marshal(newConfigMap)
	fmt.Printf("onUpdate:%s to %s\r\n", string(oldConfigMapBytes), string(newConfigMapBytes))

	oldMap := oldConfigMap.Data
	newMap := newConfigMap.Data
	diff := cmp.Diff(oldMap, newMap)
	println(diff)
}

func onDelete(obj interface{}) {
	/*configMap := obj.(*v1.ConfigMap)
	fmt.Printf("onDelete:%s\r\n", configMap.Name)
	deploymentBytes, _ := json.Marshal(configMap)
	fmt.Printf("onDelete:%s\r\n", string(deploymentBytes))*/
}
