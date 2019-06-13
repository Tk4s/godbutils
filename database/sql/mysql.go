package sql

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbs struct {
	Instances map[string]*gorm.DB
}

// DB 实例
var DB *gorm.DB

// GetMysqlInit 获取实例
func GetMysqlInit(userName, passWord, host string, port int, dbName, charset, parseTime, loc string, maxIde, maxOpen int) {
	// dbInstances := make(map[string]*gorm.DB

	// for k, v := range mysqlConfig {
	sqlURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", userName, passWord, host, port, dbName, charset, parseTime, loc)
	db, err := gorm.Open("mysql", sqlURL)
	if err != nil {
		panic(err.Error())
	}

	err = db.DB().Ping()

	if err != nil {
		panic(err.Error())
	}

	DB = db

	DB.DB().SetMaxIdleConns(maxIde)
	DB.DB().SetMaxOpenConns(maxOpen)
	// dbInstances[k] = Db
	// }

	// Dbs.Instances = dbInstances

}
