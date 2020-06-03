package cluster

import (
	"fmt"
	"testing"
)


func TestGetPodList(t *testing.T) {
	podlist := GetPodList("")
	println(podlist)
	for i,k:=range podlist{
		fmt.Println(i,k.Name,k.Namespaces,k.Containers)

	}
}
//func TestGetNameSpaces(t *testing.T) {
//	GetNameSpaces()
//}