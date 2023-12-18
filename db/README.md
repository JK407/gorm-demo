# db #

## 概述 ##
该文件夹为使用gorm-gen工具自动生成model和query文件夹

## 使用方法 ##
### 1. 配置数据源 ###
在db目录下面创建一个db.yaml文件，内容如下：
```yaml
database:
  host: 60.204.233.80
  port: 3306
  user: root
  password: Password(l4m-?9C7
  dbname: trans_demo
```
### 2. 生成model和query文件夹 ###
```shell
cd ./db
go run main.go
```
