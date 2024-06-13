package dbgormmysql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
)

type dbGorm struct{}

var O dbGorm
var I *gorm.DB

func (obj *dbGorm) Init(c *dba.DbCfg, a *appa.AppCfg) {
	if c.Dsn == "" {
		log.Fatal("Database DSN is not provided, please check DbCfg in the configuration file")
	}

	gormD := postgres.Open(c.Dsn)

	db, err := gorm.Open(gormD, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	} else {
		I = db
		log.Println("Instantiation for database-connetion using db-gorm-mysql, status: DONE!!")
	}
}
