#启动
> git clone https://github.com/FrelDX/kube-console-on-ssh.git
 && cd kube-console-on-ssh && ./k8s
 
#ssh k8s连接控制台
> rm -f ~/.ssh/known_hosts && ssh 127.0.0.1 -p 2222

#使用
>输入 p查看 pod列表
>输入 quit 退出当前程序，在容器中输入exit退出当前容器
>输入对应的数字直接登录到容器，和jumpserver一样
>默认ssh密码12345678，后续优化配置文件，自己配置
>默认ssh 端口2222，后续优化配置文件，自己配置
>大写的注意，目前支持在k8s集群外部部署，需要~/.kube/config文件认证，后续优化