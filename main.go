package main

import (
	"context"
	"fmt"
	qinkubeversioned "github.com/prettysolution/qinkube/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
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
	fmt.Println(queues)
}
