package dbgormmysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
)

type dbGorm struct{}

var O dbGorm
var I *gorm.DB
var GormConfig *gorm.Config

func (obj *dbGorm) Init(dbCfg *dba.DbCfg, appCfg *appa.AppCfg) {
	if dbCfg.Dsn == "" {
		log.Fatal("Database DSN is not provided, please check DbCfg in the configuration file")
	}

	gormD := mysql.Open(dbCfg.Dsn)

	db, err := gorm.Open(gormD, GormConfig)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		I = db
		log.Println("Instantiation for database-connetion using db-gorm-mysql, status: DONE!!")
	}
}

// some default values configuration
func init() {
	GormConfig = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	}
}
