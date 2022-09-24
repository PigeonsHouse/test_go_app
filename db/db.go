package db

import (
	"test_go_app/utils"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Psql *gorm.DB
)

func InitDB() (err error) {
	fmt.Println(utils.DbPass)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", utils.DbHost, utils.DbUser, utils.DbPass, utils.DbName, utils.DbPort)
	fmt.Println(dsn)
	Psql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	if err = Psql.AutoMigrate(); err != nil {
		return
	}
	return
}
