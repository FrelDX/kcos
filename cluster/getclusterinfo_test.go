package cluster

import (
	"fmt"
	"testing"
)


func TestGetPodList(t *testing.T) {
	podlist := GetPodList("")
	for i,k:=range podlist{
		fmt.Println(i,k.name,k.Namespaces,k.Containers)

	}
}
//func TestGetNameSpaces(t *testing.T) {
//	GetNameSpaces()
//}