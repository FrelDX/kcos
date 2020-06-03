package common

import (
	"k8s.io/client-go/kubernetes"
)

func NewClient()  (*kubernetes.Clientset){
	config:=Config()
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil
	}
	return client
}
