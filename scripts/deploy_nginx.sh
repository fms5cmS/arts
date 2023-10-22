#!/bin/bash
# author: xxx
# version: v1
# date: 2023-10-22

## 检查上一步是否执行成功
check_ok() {
    if [ $? -ne 0 ]; then
        echo "$1 error."
        exit 1
    fi
}

download_nginx() {
    cd /usr/local/src
    ## 检查文件是否存在
    if [ -f nginx-1.23.0.tar.gz ]; then
        echo "当前目录已存在 nginx-1.23.0.tar.gz"
        echo "检查 md5"
        nginx_md5=$(md5sum nginx-1.23.0.tar.gz | awk '{print $1}')
        ## 可以提前下载该文件后执行上面的 md5sum 计算得到 md5 值
        if [ "${nginx_md5}" == 'e8768e388f26fb3d56a3c88055345219' ]; then
            ## 文件存在，且 md5 值正确，则退出当前函数
            return 0
        else
            ## 文件存在，但 md5 值不正确，备份原来的文件
            sudo /bin/mv nginx-1.23.0.tar.gz nginx-1.23.0.tar.gz.old
        fi
    fi
    ## 下载文件
    sudo curl -O http://nginx.org/download/nginx-1.23.0.tar.gz
    check_ok "下载 Nginx"
}

install_nginx() {
    cd /usr/local/src
    echo "解压 Nginx"
    sudo tar zxf nginx-1.23.0.tar.gz
    check_ok "解压 Nginx"
    cd nginx-1.23.0

    echo "安装依赖"
    ## 检测有没有 yum 命令
    if which yum >/dev/null 2>&1; then
        for pkg in gcc make pcre-devel zlib-devel openssl-devel; do
            ## 查询包是否存在
            if ! rpm -1 $pkg >/dev/null 2>&1; then
                sudo yum install -y $pkg
                check_ok "yum 安装$pkg"
            else
                echo "$pkg已安装"
            fi
        done
    fi

    echo "configure Nginx"
    sudo ./configure --prefix=/usr/local/nginx --with-http_ssl_module
    check_ok "Configure Nginx"

    echo "编译和安装"
    sudo make && sudo make install
    check_ok "编译和安装"
}

start_svc() {
    echo "编辑 systemd 服务管理脚本"
    cat >/tmp/nginx.service <<EOF
[Unit]
Description=nginx - high performance web server
Documentation=http://nginx.org/en/docs/
After=network-online.target remote-fs.target nss-lookup.target
Wants=network-online.target

[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/bin/sh -c "/bin/kill -s HUP \$(/bin/cat /usr/local/nginx/logs/nginx.pid)"
ExecStop=/bin/sh -c "/bin/kill -s TERM \$(/bin/cat /usr/local/nginx/logs/nginx.pid)"

[Install]
WantedBy=multi-user.target
EOF

    sudo /bin/mv /tmp/nginx.service /lib/systemd/system/nginx.service
    check_ok "编辑 nginx.service"

    echo "加载服务"
    sudo systemctl unmask nginx.service
    sudo systemctl daemon-reload
    sudo systemctl enable nginx
    echo "启动 Nginx"
    sudo systemctl start nginx
    check_ok "启动 Nginx"
}

download_nginx
install_nginx
start_svc
