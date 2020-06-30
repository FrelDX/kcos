package common

import (
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"k8s.io/client-go/rest"
	"path/filepath"
)

func GetHome() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
// out  cluster or in cluster
func Config() (*restclient.Config) {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(GetHome(),".kube/config")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err)
		}
		return  config
	}
	return  config
}

