# 1. 安装 Docker CE（即 Docker 社区版）

使用curl下载并安装脚本

```shell
curl -sSL https://get.daocloud.io/docker | sh
```

（如果是root用户，请忽略）普通用户需要设置成非root用户也能执行docker，需要将用户加入docker组（例如你的登录用户名是admin）

```shell
sudo usermod -aG docker admin # 需要重启生效
```

# 2. Docker 更换阿里镜像源

进入 [https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors](https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors)
申请专属镜像加速器

使用 /etc/docker/daemon.json来配置 Daemon ：

```shell
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["你申请的加速器地址"]
}
EOF
```

重载daemon、重启docker

```shell
sudo systemctl daemon-reload
sudo service docker restart
```

查看docker信息

```shell
docker info
```

# 3. 安装 Docker Compose

下载Docker Compose

```shell
curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.4/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
```

配置执行权限

```shell
sudo chmod +x /usr/local/bin/docker-compose
```

检查是否安装成功

```shell
docker-compose -v
```
