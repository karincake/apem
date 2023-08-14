package databasegorm

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	lz "github.com/karincake/apem/loggerzap"
)

// Configuration type that is used by the core
type DbConf struct {
	Dsn          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxIdleTime  string `yaml:"maxIdleTime"`
	Dialect      string `yaml:"dialect"`
}

// Instance of the logger
var I *gorm.DB

var autoMigrateList []interface{}

func Init(conf DbConf) {
	if conf.Dsn == "" {
		lz.I.Warn("instantiation", zap.String("feature", "database"), zap.String("source", "gorm"), zap.String("status", "skipped"))
		return
	}
	if conf.Dialect != "mysql" && conf.Dialect != "postgres" {
		lz.I.Fatal("invalid database Dialect configuration!")
	}

	var gormD gorm.Dialector
	if conf.Dialect == "mysql" {
		gormD = mysql.Open(conf.Dsn)
	} else if conf.Dialect == "postgres" {
		gormD = postgres.Open(conf.Dsn)
	}

	db, err := gorm.Open(gormD, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		lz.I.Fatal(err.Error())
	} else {
		I = db
		lz.I.Info("instantiation", zap.String("feature", "database"), zap.String("source", "gorm"), zap.String("status", "done"))
	}

	I.AutoMigrate(autoMigrateList...)
}

// To add a model into migration list
func AutoMigrate(model ...interface{}) {
	autoMigrateList = append(autoMigrateList, model...)
}
