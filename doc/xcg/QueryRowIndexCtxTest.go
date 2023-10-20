package main

import (
	"context"
	"flag"
	"fmt"
	"looklook/app/usercenter/model"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

// 需要移动到项目更目录下执行
type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}

var configFile = flag.String("f", "app/usercenter/cmd/rpc/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c Config
	conf.MustLoad(*configFile, &c)

	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	UserModel := model.NewUserModel(sqlConn, c.Cache)

	var Mobile string = "15618918500"
	user, err := UserModel.FindOneByMobile(context.Background(), Mobile)
	if err != nil && err != model.ErrNotFound {
		fmt.Printf("Error: %v\n", err)
	}
	if user != nil {
		fmt.Printf("Error: Register user exists mobile:%s,err:%v\n", Mobile, err)
	}
}
