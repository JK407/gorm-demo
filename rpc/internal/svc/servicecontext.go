package svc

import (
	"gorm.io/gorm"
	"test_demo/common/gorm_common"
	"test_demo/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DemoDB *gorm.DB
	//DemoRedis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	demoDb, dbErr := gorm_common.InitGorm(c.Mysql.DemoDSN)
	if dbErr != nil {
		panic("连接mysql数据库失败, error=" + dbErr.Error())
	}
	//demoRdb, rdbErr := gorm_common.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	//if dbErr != nil {
	//	panic("连接redis失败, error=" + rdbErr.Error())
	//
	//}
	return &ServiceContext{
		Config: c,
		DemoDB: demoDb,
		//DemoRedis: demoRdb,
	}
}
