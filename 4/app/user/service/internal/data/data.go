package data

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"go_crouse/4/app/user/service/internal/biz"
	"go_crouse/4/app/user/service/internal/dao"
	"go_crouse/4/app/user/service/internal/data/ent"
)

var ProviderSet = wire.NewSet(NewUserData)

type UserData struct {
	db    *ent.Client
	cache *redis.Client
}

func NewUserData() biz.UserRepo {
	return &UserData{
		db:    dao.DbClient,
		cache: dao.ClientRedis,
	}
}
