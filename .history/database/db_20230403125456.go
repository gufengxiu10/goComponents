package database

import (
	"errors"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type database struct {
	db *gorm.DB
	databaseConfig
}

type databaseConfig struct {
	host     string
	name     string
	password string
	database string
	port     string
	prefix   string
}

var db *database
var once sync.Once

func init() {
	db = &database{}
}

func Instance() *gorm.DB {
	return db.db
}

func New(host, name, password, database string, args ...options) *database {
	db.databaseConfig = databaseConfig{
		host:     host,
		name:     name,
		password: password,
		database: database,
		prefix:   "",
	}

	db.With(args...)
	return db
}

func (d *database) With(args ...options) {
	for _, v := range args {
		v(db)
	}
}

func (d *database) Init() {
	once.Do(func() {
		if d.host == "" {
			panic(errors.New("mysql host is not defined"))
		}

		if d.name == "" {
			panic(errors.New("mysql username is not defined"))
		}

		if d.password == "" {
			panic(errors.New("mysql password is not defined"))
		}

		if d.database == "" {
			panic(errors.New("mysql database is no defined"))
		}

		if d.port == "" {
			d.port = definedConfig["port"]
		}

		dsn := d.name + ":" + d.password + "@tcp(" + d.host + ":" + d.port + ")/" + d.database + "?charset=utf8"
		config := mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256, // string 类型字段默认长度
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}

		client, err := gorm.Open(mysql.New(config), &gorm.Config{
			NamingStrategy: namingStrategy{
				schema.NamingStrategy{
					TablePrefix:   d.prefix,
					SingularTable: true,
					NoLowerCase:   false,
				},
			},
		})
		if err != nil {
			panic(errors.New("mysql connect failed:" + err.Error()))
		}

		client = client.Debug()
		db.db = client
	})

	fmt.Println("mysql connect succeeded")
}
