package data

import (
	"database/sql"
	"github.com/go-redis/redis"
	"go_crouse/4/app/user/service/internal/dao"
)

type UserData struct {
	db    *sql.DB
	cache *redis.Client
}

func NewUserData() *UserData {
	return &UserData{
		db:    dao.DB,
		cache: dao.ClientRedis,
	}
}
