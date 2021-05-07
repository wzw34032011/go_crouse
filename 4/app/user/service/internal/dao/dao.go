package dao

import (
	"database/sql"
	"github.com/go-redis/redis"
)

var DB *sql.DB
var ClientRedis *redis.Client

func init() {
	InitMysql(&MysqlConf{username: "sscf_admin", password: "admin@sscf_50mysql", database: "sscf_product_centre", server: "192.168.11.68", port: 3306, network: "tcp"})
	//InitMysql(&MysqlConf{})

	//initRedis(&RedisConf{})
}
