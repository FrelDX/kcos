package main

import (
	"github.com/gliderlabs/ssh"
	"k8s/pty"
	"log"
	"os"
	"syscall"
	"unsafe"
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
	log.Fatal(ssh.ListenAndServe(":2222", nil,ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
		return pass == "12345678"
	}),))
	log.Println("ssh exit")
}
func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}