package sql

import (
	"fmt"

	"github.com/pkg/errors"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbs struct {
	Instances map[string]*gorm.DB
}

type Config struct {
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	UserName  string `toml:"username"`
	Password  string `toml:"password"`
	Database  string `toml:"database"`
	MaxIde    int    `toml:"max_ide"`
	MaxOpen   int    `toml:"max_open"`
	Charset   string `toml:"charset"`
	ParseTime string `toml:"parse_time"`
	Loc       string `toml:"loc"`
}

// DB 实例
var DB *gorm.DB
var DBs = &dbs{}

func init() {
	DBs.Instances = make(map[string]*gorm.DB)
}

func InitInstanceWithName(instanceName string, cnf Config) {
	sqlURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", cnf.UserName, cnf.Password, cnf.Host, cnf.Port, cnf.Port, cnf.Charset, cnf.ParseTime, cnf.Loc)
	db, err := gorm.Open("mysql", sqlURL)
	if err != nil {
		panic(err.Error())
	}

	err = db.DB().Ping()

	if err != nil {
		panic(err.Error())
	}

	db.DB().SetMaxIdleConns(cnf.MaxIde)
	db.DB().SetMaxOpenConns(cnf.MaxOpen)

	DBs.Instances[instanceName] = db
}

func GetInstanceWithName(instanceName string) (*gorm.DB, error) {
	if db, ok := DBs.Instances[instanceName]; ok {
		if err := db.DB().Ping(); err != nil {
			return nil, err
		}
		return db, nil
	} else {
		return nil, errors.New("Instance not found.")
	}
}

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

	db.DB().SetMaxIdleConns(maxIde)
	db.DB().SetMaxOpenConns(maxOpen)

	DB = db

	// dbInstances[k] = Db
	// }

	// Dbs.Instances = dbInstances

}
