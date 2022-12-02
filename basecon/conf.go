package basecon

import (
	"fmt"
	ent "restgin/entitas"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func init() {
	conn, err := gorm.Open("postgres", "host=localhost user=golang dbname=gin_rest sslmode=disable password=22")
	if err != nil {
		fmt.Println(err)
	}
	Db = conn
	Db.Debug().AutoMigrate(&ent.User{}, &ent.Product{})

}
