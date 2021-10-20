package main

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	informer := informerFactory.Apps().V1().Deployments()
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

	deployments, _ := lister.Deployments("default").List(labels.Everything())
	for _, deployment := range deployments {
		fmt.Printf("%s\r\n", deployment.Name)
	}
	<-stopCh
}



func onAdd(obj interface{}) {
	deployment := obj.(*v1.Deployment)
	fmt.Printf("onAdd:%s\r\n", deployment.Name)
	deploymentBytes, _ := json.Marshal(deployment)
	fmt.Printf("onAdd:%s\r\n", string(deploymentBytes))
}

func onUpdate(old, new interface{}) {
	oldDeployment := old.(*v1.Deployment)
	newDeployment := new.(*v1.Deployment)
	fmt.Printf("onUpdate:%s to %s\r\n", oldDeployment.Name, newDeployment.Name)
	oldDeploymentBytes, _ := json.Marshal(oldDeployment)
	newDeploymentBytes, _ := json.Marshal(newDeployment)
	fmt.Printf("onUpdate:%s to %s\r\n", string(oldDeploymentBytes), string(newDeploymentBytes))
}

func onDelete(obj interface{}) {
	deployment := obj.(*v1.Deployment)
	fmt.Printf("onDelete:%s\r\n", deployment.Name)
	deploymentBytes, _ := json.Marshal(deployment)
	fmt.Printf("onDelete:%s\r\n", string(deploymentBytes))
}

func test(clientset *kubernetes.Clientset) {
	configMaps, _ := clientset.CoreV1().ConfigMaps("").List(metav1.ListOptions{})
	for i, cm := range configMaps.Items {
		fmt.Printf("[%d] %s\n", i, cm.GetName())
		fmt.Printf("[%d] %s\n", i, cm.Data)

		clientset.CoreV1().ConfigMaps("").Update(&cm)
	}

	sharedInformerFactory := informers.NewSharedInformerFactory(clientset, time.Minute*10)
	stopCh := make(chan struct{})
	sharedInformerFactory.Start(stopCh)

	podLister := sharedInformerFactory.Core().V1().Pods().Lister()
	podLister.List(labels.Nothing())
	podLister.Pods("kube-system").Get("kube-dns")
}

/*// NewFilteredPodInformer用来创建Pod类型的SharedIndexInformer，其中kubernetes.Interface用来实现ListerWatcher
// kubernetes.Interface是啥？但是它的实现kubernetes.Clientset应该很熟悉了吧，所以ListerWatcher就是使用kubernetes.Clientset各个资源的List和Watch函数实现的
// 这样Clientset和SharedIndexInformer之间的就联系起来了，至于函数的其他参数的说明请阅读相关的文档。
func NewFilteredPodInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			// ListFunc就是Clientset的List函数
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Pods(namespace).List(context.TODO(), options)
			},
			// WatchFunc就是Clientset的Watch函数
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Pods(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1.Pod{},
		resyncPeriod,
		indexers,
	)
}*/
