package data

import (
	"context"
	"log"
)

type UserModel struct {
	id   int
	name string
}

func (ud *UserData) AddUser(ctx context.Context, user UserModel) {
	//ud.db.Exec("")
	log.Println("data-Adduser")
}
