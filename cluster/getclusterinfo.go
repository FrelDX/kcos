package cluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/FrelDX/kcos/common"
	"log"
)

type PodList struct {
	Namespaces string
	Name       string
	Containers []string
	Ip  string
}

func GetPodList(namespace string)([]PodList,error) {
	var pod []PodList
	pods, err := common.NewClient().CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	podtmp := PodList{}
	for i := 0; i < len(pods.Items); i++ {
		podtmp.Namespaces = pods.Items[i].Namespace
		podtmp.Ip = pods.Items[i].Status.PodIP
		podtmp.Name = pods.Items[i].Name
		Containers := []string{}
		// get Containers to pod info
		for p := 0; p < len(pods.Items[i].Spec.Containers); p++ {
			Containers = append(Containers, pods.Items[i].Spec.Containers[p].Name)
		}
		podtmp.Containers = Containers
		pod = append(pod, podtmp)
	}
	return pod,err
}
func GetNameSpaces()([]string,error){
	name :=[]string{}
	Namespaces, err := common.NewClient().CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Print(err)
		return name,err
	}
	for i := 0; i < len(Namespaces.Items); i++ {
		name = append(name,Namespaces.Items[i].Name)
	}
	return name,nil
}