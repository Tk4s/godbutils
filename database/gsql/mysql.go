package gsql

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type instance struct {
	DB *gorm.DB
}

// DB 实例
var Instance = &instance{}


// NewInstance 获取实例
func NewInstance(userName, passWord, host string, port int, dbName, charset, parseTime, loc string, maxIde, maxOpen int) error {
	// dbInstances := make(map[string]*gorm.DB

	// for k, v := range mysqlConfig {
	sqlURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", userName, passWord, host, port, dbName, charset, parseTime, loc)
	db, err := gorm.Open("mysql", sqlURL)
	if err != nil {
		return err
	}

	err = db.DB().Ping()

	if err != nil {
		return err
	}



	db.DB().SetMaxIdleConns(maxIde)
	db.DB().SetMaxOpenConns(maxOpen)

	Instance.DB = db
	return nil
}


