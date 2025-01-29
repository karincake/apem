package msredis

import (
	"log"

	"github.com/go-redis/redis"

	"github.com/karincake/apem/appa"
	lo "github.com/karincake/apem/loggero"
	"github.com/karincake/apem/msa"
)

type msRedis struct{}

var O msRedis
var I *redis.Client // instance

func (obj *msRedis) Init(conf *msa.MsCfg, app *appa.AppCfg) {
	if len(conf.Dsn) == 0 {
		log.Fatal("Instantiation for memory storage using redis failed: no dsn provided at the 'config.yml' file")
		return
	}

	I = redis.NewClient(&redis.Options{
		Addr: conf.Dsn,
	})
	_, err := I.Ping().Result()
	if err != nil {
		panic(err)
	}
	lo.I.Println("Instantiation for memory storage using redis, status: DONE!!")
}
