package main

import (
	"context"
	"entgo.io/ent/dialect"
	"go_crouse/ent_test/ent"
	"go_crouse/ent_test/ent/user"
	"log"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open(dialect.MySQL, "user:password@tcp(127.0.0.1:3306)/database?parseTime=True")
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	/*u, err := client.User.Create().
		SetAge(30).
		SetName("a8m").
		SetGender(user.Gender0).
		Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("user was created: ", u)*/
	users, error := client.User.Query().Order(ent.Desc(user.FieldAge)).All(context.Background())
	if error != nil {
		log.Fatal(error)
	}
	log.Println(users)
}
