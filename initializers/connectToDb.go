package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func connectToDb() {
	var err error

	dsn := os.Getenv("DATABASE_USER") + ":@tcp(127.0.0.1:3306)/go-auth?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error Connecting to database")
	}

}
