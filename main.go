package main

import (
	"github.com/gliderlabs/ssh"
	"kube-console-on-ssh/pty"
	"log"
)
func main()  {
	sshd()
}
func sshd()  {
	ssh.Handle(func(s ssh.Session) {
		//cmd := exec.Command("top")
		//ptyReq, winCh, isPty := s.Pty()
		pty.MainInterface(s)
	})
	log.Println("starting ssh server on port 2222...")
	ssh.ListenAndServe(":2222", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
		return pass == "12345678"}),
		ssh.HostKeyFile("./id_rsa"),
		)
	log.Println("ssh exit")
}

