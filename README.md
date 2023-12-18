# Gorm-Demo #
该文档用于介绍gorm如何在go-zero中使用。

## 概述 ##
项目环境：
- go version go1.20.2 windows/amd64
- goctl version 1.6.0
- protoc version 3.20.3
- protoc-gen-go version 1.31.0
- protoc-gen-go-grpc version 1.3.0 

该项目是一个使用go-zero框架并将gorm集成生成的事务Demo项目，用于演示go-zero框架集成gorm事务使用方法。
db文件夹下面的model中有三张表，分别为用户表、商品表和订单表， 主要场景为用户下单，下单时会同时更新商品库存表和订单表，如果任何一个步骤失败，整个事务将会回滚，确保数据的一致性和完整性。
该项目提供了两个接口，分别为单用户下单和多用户下单，其中单用户下单接口支持下单和取消订单两种操作，多用户下单接口和取消订单两种操作。
这两个接口的实现都是通过事务来实现的，事务的实现方式有两种，一种是通过乐观锁实现，另一种是通过悲观锁实现。
在单用户下单接口中，会启动一个事务，事务将包含将订单插入到订单表，同时更新商品库存表的操作。如果任何一个步骤失败，整个事务将会回滚，确保数据的一致性和完整性。

## 启动服务 ##
### 启动api服务 ###
```shell
cd gorm-demo/api
go run orderapiservice.go
```

## api接口 ##
### 单用户下单 /order/user - POST action=1 下单, action=2 取消订单 ###
```shell
curl --location '127.0.0.1:8888/order/user' \
--header 'Content-Type: application/json' \
--data '{
"user_id": 1,
"product_id": 2,
"order_amount": 10,
"order_id":10,
"action": 1
}'
```

### 多用户下单 /order/multi - POST (未实现) ###
```json
{
    "user_id": [1,2],
    "product_id": 1,
    "order_amount": 1,
    "action": 1 
}
```