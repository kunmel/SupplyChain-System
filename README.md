# 部署方式

---

> 🚀基于区块链的钢材供应金融链。本项目使用Hyperledger Fabric构建区块链网络, go编写智能合约，应用层使用gin+fabric-sdk-go调用合约。前端展示使用vue+element。
## 技术栈

- Hyperledger Fabric
- Go
- Vue
- Docker

## 前提

Linux 或者 Mac，连接网络，要求安装了 Docker 和 Docker Compose 以及 Go 环境

附 Linux 安装 Docker 和 Docker Compose 教程：[点此跳转](/Install.md)

附 Linux 安装 Go 教程: https://blog.csdn.net/u014454538/article/details/88649963

## 运行

### 1、克隆本项目放在任意目录下，例：`./Desktop/blockchain-real-estate`


### 2、给予项目权限，执行 `sudo chmod -R +x ./Desktop/blockchain-real-estate/`

### 3、进入 `deploy` 目录，执行 `./start.sh` 启动区块链网络

### 4、进入 `Supplier-Vue` 目录，执行 `./build.sh` 编译前端

### 5、进入 `application` 目录，执行 `./build.sh` 编译后端

### 6、在 `application` 目录下，执行 `./start.sh` 启动应用

### 7、浏览器访问 [http://localhost:8000/web](http://localhost:8000/web)

## 目录结构

`application` : go gin + fabric-sdk-go 调用链码，提供外部访问接口，前端编译后静态资源放在`dist`目录下

`chaincode` : go 编写的智能合约

`deploy` : 区块链网络配置

`vue` : vue + element的前端展示页面
