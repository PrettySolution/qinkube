package main

import (
	"context"
	"fmt"
	qinkubeversioned "github.com/prettysolution/qinkube/pkg/client/clientset/versioned"
	qinkubeinformerfactory "github.com/prettysolution/qinkube/pkg/client/informers/externalversions"
	"github.com/prettysolution/qinkube/pkg/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		fmt.Println("Could not load kubeconfig.ClientConfig()")
		panic(err)
	}
	clientset := qinkubeversioned.NewForConfigOrDie(config)

	queues, err := clientset.QinkubeV1alpha1().Queues("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("lenght of Queues is %d\n", len(queues.Items))
	//fmt.Printf("name of first queue is '%s'\n", queues.Items[0].Name)

	informer := qinkubeinformerfactory.NewSharedInformerFactory(clientset, 20*time.Minute)
	ch := make(chan struct{})
	c := controller.NewController(clientset, informer.Qinkube().V1alpha1().Queues())

	informer.Start(ch)
	if err := c.Run(ch); err != nil {
		log.Println("error running controller", err.Error())
	}
}
