package cluster

import (
	"fmt"
	"k8s/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodList struct {
	Namespaces string
	Name       string
	Containers []string
}

func GetPodList(namespace string) []PodList {
	var pod []PodList
	pods, err := common.NewClient().CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil
	}
	podtmp := PodList{}
	for i := 0; i < len(pods.Items); i++ {
		podtmp.Namespaces = pods.Items[i].Namespace
		podtmp.Name = pods.Items[i].Name
		Containers := []string{}
		// get Containers to pod info
		for p := 0; p < len(pods.Items[i].Spec.Containers); p++ {
			Containers = append(Containers, pods.Items[i].Spec.Containers[p].Name)
		}
		podtmp.Containers = Containers
		pod = append(pod, podtmp)
	}
	return pod
}
func GetNameSpaces() error {
	Namespaces, err := common.NewClient().CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return nil
	}
	for i := 0; i < len(Namespaces.Items); i++ {
		fmt.Println(Namespaces.Items[i].Name)
	}
	return nil
}
