package data

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"go_crouse/4/app/user/service/internal/biz"
	"go_crouse/4/app/user/service/internal/dao"
)

var ProviderSet = wire.NewSet(NewUserData)

type UserData struct {
	db    *sql.DB
	cache *redis.Client
}

func NewUserData() biz.UserRepo {
	return &UserData{
		db:    dao.DB,
		cache: dao.ClientRedis,
	}
}
