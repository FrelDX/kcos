package pty

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"github.com/FrelDX/kcos/cluster"
	"github.com/FrelDX/kcos/common"
	interrupt "github.com/FrelDX/kcos/util"
	"log"
	"strconv"
)

const (
	// Fill in spaces when less than 60 characters
	DisplayLengthPod = 60
	DisplayLengthNameSpace = 20
)

// Global pod information storage, used to connect to pod shell
var err error
var n string
func (p *PtyTerminal)newPty(namespace,podName,container string) error {
	// to exec,but /bin/bash first then /bin/sh
	exec := func()  error{
		config := common.Config()
		client := common.NewClient()
		err := Remotepty(client, config, namespace, podName, "/bin/bash", container, p.Handler, p.Handler, p.Handler)
		if err != nil {
			err = Remotepty(client, config, namespace, podName, "/bin/sh", container, p.Handler, p.Handler, p.Handler)
			if err != nil {
				log.Print(err)
			}
		}
		return  nil
	}
	// Processing signal
	return interrupt.Chain(nil, func() {
		log.Println(p.User + "  :","返回主界面")
	}).Run(exec)
}

func (p *PtyTerminal)DisplayAllPod()  {
	 p.podIndex,err=cluster.GetPodList("")
	 if err !=nil{
	 	log.Println(p.User," :",err)
		 return
	 }
	// to DisplayPod
	p.DisplayPod()
}
func (p *PtyTerminal)DisplayNamespacePod()  {
	pod,err :=cluster.GetPodList(p.namespace)
	if err !=nil{
		log.Println(p.User," :",err)
		return
	}
	// to DisplayPod
	p.podIndex = pod
	p.DisplayPod()
}


func (p *PtyTerminal)DisplayNameSpace()([]string,error  ){
	namespace,err :=cluster.GetNameSpaces()
	if err!=nil{
		return namespace, err
	}
	for i,k:=range namespace{
		io.WriteString(p.Handler, fmt.Sprint(SetColorBlue(strconv.Itoa(i)),"\t",SetColorGreen(k),"\n"))
	}
	return namespace,nil
}
type  PtyTerminal struct {
	Handler ssh.Session
	Terminal  *terminal.Terminal
	User string
	podIndex []cluster.PodList
	namespace string
}

func NewPtyTerminal(s ssh.Session)  *PtyTerminal {
	return &PtyTerminal{
		Terminal:terminal.NewTerminal(s, s.User() + "# "),
		User: s.User(),
		Handler: s,
	}
}

func (p *PtyTerminal)Start()  {
	log.Println(p.User,": 登陆")
	p.WelcomePage()
	p.MainInterface()
}
func (p *PtyTerminal)stop()  {
	log.Println(p.User,":退出")
	p.Handler.Close()
}

func (p *PtyTerminal)MainInterface() {
	for {
		line, err := p.Terminal.ReadLine()
//
Restart:
		if err != nil {
			break
			log.Println(p.User, " :", err)
		}
		switch line {
		case "quit":
			p.stop()
		case "p":
			p.DisplayAllPod()
		case "n":
			namespace,err :=p.DisplayNameSpace()
			if err!=nil{
				log.Println(err)
				break
			}
			// Operations corresponding to namespace
			for{
				n, _ =  p.Terminal.ReadLine()
				if n == "m"{
					p.WelcomePage()
					break
				}
				if n == "quit"{
					p.stop()
				}
				number, err := strconv.Atoi(n)
				if err ==nil{
					// Prevent index out of range
					if number < len((namespace)){
						p.namespace = namespace[number]
						p.DisplayNamespacePod()
						break
					}else {
						io.WriteString(p.Handler,SetColorRed("No such option, please re-enter\n"))
						log.Println(err)
						continue
					}
				}
				line = n
				goto Restart
			}
		case "m":
			p.WelcomePage()
		default:
			number, err := strconv.Atoi(line)
			if err != nil{
				io.WriteString(p.Handler,SetColorRed("No such option, please re-enter\n"))
				continue
			}
			// go to pod shell
			if number < len((p.podIndex)){
				// Multiple containers
				if len(p.podIndex[number].Containers) >1{
					io.WriteString(p.Handler, fmt.Sprint(SetColorBlue("Please select a container "),"\n"))
					for i,c:= range   p.podIndex[number].Containers{
						io.WriteString(p.Handler,fmt.Sprint(SetColorBlue(strconv.Itoa(i)),"\t",SetColorRed(c),"\n"))
					}
					// Get user selected container
					container, _ := p.Terminal.ReadLine()
					containerNumber, err :=strconv.Atoi(container)
					if err ==nil{
						if containerNumber < len((p.podIndex[number].Containers)) {
							p.newPty(p.podIndex[number].Namespaces,p.podIndex[number].Name,p.podIndex[number].Containers[containerNumber])
						}
					}
				}
				p.newPty(p.podIndex[number].Namespaces,p.podIndex[number].Name,"")
			}
		}

	}
}

func (p *PtyTerminal)DisplayPod()  {
	for i,k:=range p.podIndex{
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
		io.WriteString(p.Terminal, fmt.Sprint(SetColorBlue(strconv.Itoa(i)),"\t",SetColorRed(namespace),"\t",SetColorGreen(pod),"\t",SetColorYellow(k.Ip),"\n"))
	}
}


func (p *PtyTerminal)WelcomePage()  {
	io.WriteString(p.Terminal, SetColorGreen("\t\t\t欢迎登陆kcos (kube-console-on-ssh)\n"))
	io.WriteString(p.Terminal, SetColorGreen("\t\t\t输入quit退出当前终端\n"))
	io.WriteString(p.Terminal, SetColorGreen("\t\t\t当前登陆的用户:" + p.User+"\n"))
	io.WriteString(p.Terminal, SetColorGreen("\t\t\t选择对应的数字连接到对应的pod shell\n"))
	io.WriteString(p.Terminal, SetColorGreen("\t\t\t输入p查看所有可以登陆的pod列表\n"))
	io.WriteString(p.Terminal, SetColorGreen("\t\t\t输入n 查看namespace下所有的pod\n"))
	io.WriteString(p.Terminal, SetColorGreen("\t\t\t输入m 返回到主菜单\n"))
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