# ✨ Features

- 🌈 Connect the pod in kubernetes cluster through the terminal to send instructions
- 🐿️ The SSH protocol encapsulated by golang can run on both Linux and windows
- 🚀 It can run anywhere in the kubernetes cluster with the help of yaml files

# 🐱 To Do

🛡 Perfect log audit function

♾ Perfect API interface

👬 Multi user system

# 💡what is Kube console on SSH

> Kube console on SSH connects to kubernetes pod based on SSH protocol and does not rely on kubectl Exec command, because Kube console on SSH implements the service end of SSH, so you can use SSH command to connect to Kube console on SSH on any host. If you need to connect to pod on a large scale, you don't need to install kubectl everywhere to enter the container. Through SSH protocol, it is more convenient to enter the container, and more operations will be supported in the future


# 📦  How to install

```
//If you don't have a go locale, you can run binaries directly in the repository, which is convenient
//Download code
git clone https://github.com/FrelDX/kcos.git
// Run kube-console-on-ssh
cd kcos && ./kcos
```
# ♾ Connection to kube-console-on-ssh

Using SSH client to connect to Kube console on SSH
>rm -f ~/.ssh/known_hosts && ssh 127.0.0.1 -p 2222



# 💻 How to use it

- Enter p to view the list of pods

- Enter quit to exit the program

- Enter the corresponding number into the corresponding container

- The default SSH password is 12345678, and the account can be any

![Large GIF (942x608)](./doc/readme.gif)