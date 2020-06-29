package common

import (
	restclient "k8s.io/client-go/rest"
	"os"
	"k8s.io/client-go/rest"
)

func GetHome() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
// out  cluster
func Config() (*restclient.Config) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return  config
}

