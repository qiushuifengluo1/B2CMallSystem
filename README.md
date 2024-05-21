## B2C Mall System

## 项目简介

这是一个基于B2C模式的电子商务系统，后端使用了Beego框架，前端使用了HTML5、CSS和jQuery，数据库采用了Redis和MySQL。


## 环境要求

- Ubuntu 20.04 LTS
- Go 1.16+
- MySQL 8.0+
- Redis 6.0+
- Node.js 14.0+
- Git

## 安装步骤

### 1. 安装Go

在Ubuntu上安装Go：

sudo apt update
sudo apt install -y golang-go

### 2. 安装MySQL
在Ubuntu上安装MySQL：
sudo apt update
sudo apt install -y mysql-server
sudo systemctl start mysql
sudo systemctl enable mysql

### 3.配置MySQL：
sudo mysql_secure_installation

### 4.创建数据库和用户：
CREATE DATABASE shop;
CREATE USER 'your_name'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON shop.* TO 'your_name'@'localhost';
FLUSH PRIVILEGES;

### 5. 安装Redis
在Ubuntu上安装Redis：
sudo apt update
sudo apt install -y redis-server
sudo systemctl start redis-server
sudo systemctl enable redis-server

### 6. 安装Node.js
在Ubuntu上安装Node.js：
curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash -
sudo apt install -y nodejs


### 7. 配置项目
编辑项目的配置文件 conf/app.conf：

appname = b2c-ecommerce   ----你自己项目的名字
httpport = 8080        ------自己设置端口
runmode = dev

[mysql]
user = 你数据库用户
password = 你数据库用户对应的密码
host = 127.0.0.1
port = 3306
dbname = shop

[redis]
host = 127.0.0.1
port = 6379


### 8. 安装项目依赖
在项目根目录下安装Go依赖：

go mod tidy

### 9. 运行项目
在项目根目录下运行Beego项目：
bee run

访问 http://localhost:[你上面设置的端口]
