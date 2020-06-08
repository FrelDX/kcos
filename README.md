#what is Kube console on SSH
> Kube console on SSH connects to kubernetes pod based on SSH protocol and does not rely on kubectl Exec command, because Kube console on SSH implements the service end of SSH, so you can use SSH command to connect to Kube console on SSH on any host. If you need to connect to pod on a large scale, you don't need to install kubectl everywhere to enter the container. Through SSH protocol, it is more convenient to enter the container, and more operations will be supported in the future

#how to install
```
//If you don't have a go locale, you can run binaries directly in the repository, which is convenient
//Download code
git clone https://github.com/FrelDX/kube-console-on-ssh.git
// Run kube-console-on-ssh
cd kube-console-on-ssh && ./kube-console-on-ssh
```


#connection to kube-console-on-ssh
Using SSH client to connect to Kube console on SSH
>rm -f ~/.ssh/known_hosts && ssh 127.0.0.1 -p 2222



#How to use it
1. 输入p查看pod列表
2. 输入quit退出程序
3. 输入对应的数字进入对应的容器
4. 默认的ssh密码是12345678，账号可以是任意的
