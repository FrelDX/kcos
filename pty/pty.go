package pty

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"kube-console-on-ssh/cluster"
	"kube-console-on-ssh/common"
	interrupt "kube-console-on-ssh/util"
	"log"
	"strconv"
)

const (
	// 不足60字符的时候空格补齐
	// Fill in spaces when less than 60 characters
	DisplayLengthPod = 60
	DisplayLengthNameSpace = 20
)

// 全局pod信息保存处,用于连接到pod shell
var podIndex *[]cluster.PodList
func New(namespace string,podName string,stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	// to exec,but /bin/bash first then /bin/sh
	exec := func()  error{
		config := common.Config()
		client := common.NewClient()
		err := Remotepty(client, config, namespace, podName, "/bin/bash", "", stdin, stdout, stderr)
		if err != nil {
			err = Remotepty(client, config, namespace, podName, "/bin/sh", "", stdin, stdout, stderr)
			if err != nil {
				log.Print(err)
			}
		}
		return  nil
	}
	// Processing signal
	return interrupt.Chain(nil, func() {
		log.Print("go to interface")
	}).Run(exec)
	}
func MainInterface(s ssh.Session){
	WelcomePage(s)
	term := terminal.NewTerminal(s, SetColorRed(s.User() + "#"))
	line := ""
	// get user input
	for {
		line, _ = term.ReadLine()
		if line == "quit" {
			break
		}else if line == "p" {
			DisplayAllPod(s)
		}else if line =="m"{
			WelcomePage(s)
		} else if line =="n"{
			namespace :=DisplayNameSpace(s)
			// 对应namespace的操作
			// Operations corresponding to namespace
			n:=""
			for{
				n, _ = term.ReadLine()
				if n == "m"{
					WelcomePage(s)
					break
				}
				if n == "quit"{
					break
				}
				number, err := strconv.Atoi(n)
				if err ==nil{
					// 防止索引超出范围
					if number < len((*namespace)){
						DisplayNamespacePod(s,(*namespace)[number])
					}
					log.Println("输入的索引超出范围")
				}
				break
			}
		} else if line == "d"{
			DisplayDeploy(s)
		}
		number, err := strconv.Atoi(line)
		if err == nil{
			if number < len((*podIndex)){
				New((*podIndex)[number].Namespaces,(*podIndex)[number].Name,s,s,s)
			}
			log.Println("输入的索引超出范围")
			continue
		}
	}
}
func DisplayAllPod(s ssh.Session)  {
	pod :=cluster.GetPodList("")
	// to DisplayPod
	DisplayPod(pod,s)
}

func DisplayNamespacePod(s ssh.Session,namespace string)  {
	pod :=cluster.GetPodList(namespace)
	// to DisplayPod
	DisplayPod(pod,s)
}

func SetColorGreen(msg string) string {
	return  fmt.Sprintf("\033[32;1m%s\033[0m",msg)
}
func SetColorRed(msg string) string {
	return  fmt.Sprintf("\033[31;1m%s\033[0m",msg)
}
func SetColorBlue(msg string) string {
	return  fmt.Sprintf("\033[34;1m%s\033[0m",msg)
}
func SetColorYellow(msg string) string {
	return  fmt.Sprintf("\033[33;1m%s\033[0m",msg)
}
func WelcomePage(s ssh.Session)  {
	io.WriteString(s, SetColorGreen("\t\t\t欢迎登陆kcos (kube-console-on-ssh)\n"))
	io.WriteString(s, SetColorGreen("\t\t\t输入quit退出当前终端\n"))
	io.WriteString(s, SetColorGreen("\t\t\t当前登陆的用户:" + s.User())+"\n")
	io.WriteString(s, SetColorGreen("\t\t\t选择对应的数字连接到对应的pod shell\n"))
	io.WriteString(s, SetColorGreen("\t\t\t输入p查看所有可以登陆的pod列表\n"))
	io.WriteString(s, SetColorGreen("\t\t\t输入n 查看namespace下所有的pod\n"))
	io.WriteString(s, SetColorGreen("\t\t\t输入m 返回到主菜单\n"))
}
func DisplayPod(pod *[]cluster.PodList,s ssh.Session)  {
	for i,k:=range *pod{
		//把名字的长度统一，防止显示的时候乱码
		// Unify the length of the name to prevent garbled characters when displaying
		namespace :=k.Namespaces
		if len(k.Namespaces)<DisplayLengthNameSpace{
			for i:=0;i<DisplayLengthNameSpace-len(k.Namespaces);i++{
				namespace = namespace + " "
			}
		}
		pod := k.Name
		if len(k.Name) < DisplayLengthPod{
			for i:=0;i<DisplayLengthPod-len(k.Name);i++{
				pod = pod +" "
			}
		}
		io.WriteString(s, fmt.Sprint(SetColorBlue(strconv.Itoa(i)),"\t",SetColorRed(namespace),"\t",SetColorGreen(pod),"\t",SetColorYellow(k.Ip),"\n"))
	}
	// 最新的信息刷新到全局pod信息保存处
	podIndex = pod
}

func DisplayNameSpace(s ssh.Session) *[]string  {
	namespace :=cluster.GetNameSpaces()
	for i,k:=range *namespace{
		io.WriteString(s, fmt.Sprint(SetColorBlue(strconv.Itoa(i)),"\t",SetColorGreen(k),"\n"))
	}
	return namespace
}

func DisplayDeploy(s ssh.Session)  {

}