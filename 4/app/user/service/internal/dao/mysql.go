package dao

import (
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_crouse/4/app/user/service/internal/data/ent"
	"log"
)

type MysqlConf struct {
	username string
	password string
	database string
	server   string
	port     int
	network  string
}

func InitMyClient(c *MysqlConf) {
	client, err := ent.Open(dialect.MySQL, fmt.Sprintf("%s:%s@%s(%s:%d)/%s", c.username, c.password, c.network, c.server, c.port, c.database))
	if err != nil {
		log.Fatalln(err.Error())
	}
	DbClient = client

	//初始化user表
	/*err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}*/
}
