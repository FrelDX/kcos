package pty

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"k8s/cluster"
	interrupt "k8s/util"
	"strconv"

	//"fmt"
	"k8s/common"
	//interrupt "k8s/util"
	"log"
	//"os"
	"io"
)
func New(namespace string,podName string,stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	// to exec,but /bin/bash first then /bin/sh
	exec := func()  error{
		config := common.Config()
		client := common.NewClient()
		err := Remotepty(client, config, namespace, podName, "/bin/bash", "", stdin, stdout, stderr)
		if err != nil {
			err = Remotepty(client, config, namespace, podName, "/bin/sh", "", stdin, stdout, stderr)
			if err != nil {
				log.Fatal(err)
			}
		}
		return  nil
	}
	// Processing signal
	return interrupt.Chain(nil, func() {
		// print
		println("go to interface")
	}).Run(exec)
	}
func MainInterface(s ssh.Session){
	io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	io.WriteString(s, fmt.Sprint("number\t\t\t","namespace","\t\t\t","podname\n"))
	//get pod list
	pod :=cluster.GetPodList("")
	for i,k:= range pod{
		io.WriteString(s, fmt.Sprint(i,"\t",k.Namespaces,"\t\t\t",k.Name,"\n"))
	}
	term := terminal.NewTerminal(s, ">")
	line := ""
	// get user input
	for {
		line, _ = term.ReadLine()
		if line == "quit" {
			break
		}
		if line == "p" {
			Console(s)
			continue
		}
		//Number of users selected
		number, err := strconv.Atoi(line)
		if err ==nil{
			// go to pod shell
			New(pod[number].Namespaces,pod[number].Name,s,s,s)

		}
		io.WriteString(s, fmt.Sprintln("please enter a number"))
	}
}
	//New("default","fast",s,s,s)
func Console(s ssh.Session)  {
	pod :=cluster.GetPodList("")
	for i,k:= range pod{
		io.WriteString(s, fmt.Sprint(i,"\t",k.Namespaces,"\t\t\t",k.Name,"\n"))
	}
}