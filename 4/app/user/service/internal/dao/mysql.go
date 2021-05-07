package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func InitMysql(c *MysqlConf) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", c.username, c.password, c.network, c.server, c.port, c.database)
	MysqlDb, error := sql.Open("mysql", dsn)
	if error != nil {
		log.Fatalln(error.Error())
	}
	DB = MysqlDb

	/*error = DB.Ping()
	if error != nil{
		log.Fatalln(error.Error())
	}*/
}
