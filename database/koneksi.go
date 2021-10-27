package database

import (
	"fmt"
	"spewebinar/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := config.Config("DB_USER") + ":" + config.Config("DB_PASS") + "@(" + config.Config("DB_HOST") + ":" + config.Config("DB_PORT") + ")/" + config.Config("DB_NAME")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil{
		return nil, err
	}
	fmt.Println("berhasil koneksi ke database")
	return db, nil
}
