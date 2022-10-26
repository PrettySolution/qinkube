package controller

import (
	qinkubeversioned "github.com/prettysolution/qinkube/pkg/client/clientset/versioned"
	qinkubeinformer "github.com/prettysolution/qinkube/pkg/client/informers/externalversions/qinkube/v1alpha1"
	qinkubelister "github.com/prettysolution/qinkube/pkg/client/listers/qinkube/v1alpha1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"log"
	"time"
)

type Controller struct {
	clientset     qinkubeversioned.Interface
	qinkubeSynced cache.InformerSynced
	qinkubeLister qinkubelister.QueueLister
	wq            workqueue.RateLimitingInterface
}

func NewController(clientset qinkubeversioned.Interface, informer qinkubeinformer.QueueInformer) *Controller {
	c := &Controller{
		clientset:     clientset,
		qinkubeSynced: informer.Informer().HasSynced,
		qinkubeLister: informer.Lister(),
		wq:            workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
	}

	informer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,
			DeleteFunc: c.handleDel,
		})

	return c
}

func (c *Controller) handleAdd(obj interface{}) {
	log.Println("handleAdd was called")
	c.wq.Add(obj)
}

func (c *Controller) handleDel(obj interface{}) {
	log.Println("handleDel was called")
	c.wq.Forget(obj)
}

func (c *Controller) Run(ch chan struct{}) error {
	if ok := cache.WaitForCacheSync(ch, c.qinkubeSynced); !ok {
		log.Println("cache was not synced")
	}

	go wait.Until(c.worker, time.Second, ch)

	<-ch
	return nil
}

func (c *Controller) worker() {
	for c.processNexItem() {

	}
}

func (c *Controller) processNexItem() bool {
	item, shutDown := c.wq.Get()
	if shutDown {
		return false
	}

	//Todo: forget item one has been processed.
	defer c.wq.Forget(item)
	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		log.Printf("error %s calling NamespaceKeyFunc on cache for item", err.Error())
	}

	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		log.Printf("error splitting key into namespace and name %s", err.Error())
		return false
	}

	queue, err := c.qinkubeLister.Queues(ns).Get(name)
	if err != nil {
		log.Printf("error getting the Queue from lister %s", err.Error())
		return false
	}

	log.Println(queue.Spec)

	return true
}
