package dao

import (
	"github.com/go-redis/redis"
	"go_crouse/4/app/user/service/internal/data/ent"
)

var DbClient *ent.Client
var ClientRedis *redis.Client

func init() {
	//InitMyClient(&MysqlConf{username: "sscf_admin", password: "admin@sscf_50mysql", database: "sscf_product_centre", server: "192.168.11.68", port: 3306, network: "tcp"})
	//InitMyClient(&MysqlConf{})

	//initRedis(&RedisConf{})
}
