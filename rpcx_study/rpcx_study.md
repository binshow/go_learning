安装 rpcx 的基础功能：

```zsh
go get -u -v -tags "reuseport quic kcp zookeeper etcd consul ping" github.com/smallnest/rpcx/...
```

tags 对应:

quic: 支持 quic 协议
kcp: 支持 kcp 协议
zookeeper: 支持 zookeeper 注册中心
etcd: 支持 etcd 注册中心
consul: 支持 consul 注册中心
ping: 支持 网络质量负载均衡
reuseport: 支持 reuseport

