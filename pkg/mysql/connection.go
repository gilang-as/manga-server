package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func (d *DatabaseClient) WithTransaction(fc func(tx *gorm.DB) error) (err error){
	return d.DB.Transaction(fc)
}

func Connect(dbHost string, user string, password string, databaseName string) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local&net_write_timeout=6000", user, password, dbHost, databaseName))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return db
}