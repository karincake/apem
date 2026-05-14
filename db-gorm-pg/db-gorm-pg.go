package dbgormmysql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
	lo "github.com/karincake/apem/loggero"
)

type dbGorm struct{}

var O dbGorm
var I *gorm.DB
var IS map[string]*gorm.DB
var GormConfig *gorm.Config

func init() {
	IS = make(map[string]*gorm.DB)
}

func (obj *dbGorm) Init(dbCfg *dba.DbCfg, a *appa.AppCfg) {
	if dbCfg.Dsn == "" {
		log.Fatal("Database DSN is not provided, please check DbCfg in the configuration file")
	}

	gormD := postgres.Open(dbCfg.Dsn)

	db, err := gorm.Open(gormD, GormConfig)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		I = db
		lo.I.Println("Instantiation for database-connetion using db-gorm-pg, status: DONE!!")
	}
}

func (obj *dbGorm) InitMulti(dbCfgs *dba.MultiDbCfg, a *appa.AppCfg) {
	for _, dbCfg := range dbCfgs.Dbs {
		if dbCfg.Dsn == "" {
			log.Fatal("Database with name: " + dbCfg.Name + " DSN is not provided, please check MultiDbCfg in the configuration file")
		}

		gormD := postgres.Open(dbCfg.Dsn)

		db, err := gorm.Open(gormD, GormConfig)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			lo.I.Println("Instantiation for database-connetion using db-gorm-pg with name: " + dbCfg.Name + ", status: DONE!!")
		}

		IS[dbCfg.Name] = db
	}
}
