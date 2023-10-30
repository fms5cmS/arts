
学习视频：[Shell 脚本一天一练](https://www.bilibili.com/video/BV1Cy4y1A7UC)

[源代码](https://github.com/aminglinux/daily_shell/tree/main)

# Redis 源码部署注意

源码安装时注意：

```shell
# 安装库
yum -y install gcc zlib zlib-devel pcre-devel openssl openssl-devel
# 源码安装 redis 时
make MALLOC=l  ibc

# 修改 redis 配置
daemonize yes
bind 0.0.0.0 #任意ip都可以连接
protected-mode no #关闭保护，允许非本地连接
port 6379 #端口号
dir /usr/loacl/redis-6.0.10/data/ #db等相关目录位置
logfile "/usr/local/redis-6.0.10/logs/redis_6379.log"
appendonly yes #开启日志形式
```