package orm

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
}

func TestOrm() {
	db, err := gorm.Open("mysql", "root:1234567@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var user User
	err = db.Where("id = ?", 1).First(&user).Error
	log.Printf("user: %+v, error: %s", user, err)
}
