package basecon

import (
	"fmt"
	ent "restgin/entitas"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func init() {
	conn, err := gorm.Open("postgres", "host=nama_host user=nama_user_golang dbname=nama_database sslmode=disable password=pass_postgres")
	if err != nil {
		fmt.Println(err)
	}
	Db = conn
	Db.Debug().AutoMigrate(&ent.User{}, &ent.Product{})

}
