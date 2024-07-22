package configs

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:Password@tcp(127.0.0.1:3306)/ecomm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	fmt.Println("coba")
	fmt.Println("coba")
	DB = db
	return DB, nil
}
