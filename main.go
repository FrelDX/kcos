package main

import (
	"github.com/FrelDX/kcos/pty"
	"github.com/gliderlabs/ssh"
	"log"
)
func main()  {

	sshd()

}
func sshd()  {
	ssh.Handle(func(s ssh.Session) {
		//pty.MainInterface(s)
		ss :=pty.NewPtyTerminal(s)
		ss.Start()
	})
	log.Println("starting ssh server on port 2222...")
	log.Fatal(
		ssh.ListenAndServe(":2222", nil,
			ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
				return pass == "12345678"}),
			ssh.HostKeyFile("./key/id_rsa"),
		),
	)
}