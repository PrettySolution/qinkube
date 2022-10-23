package main

import (
	"fmt"

	"k8s.io/vfs-workflows/pkg/apis/workflow/v1alpha1"
)

func main() {
	k := v1alpha1.Kluster{}
	fmt.Println(k)
}
