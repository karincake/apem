package msredis

import (
	"log"

	"github.com/go-redis/redis"

	"github.com/karincake/apem/msa"
)

type msRedis struct{}

var O msRedis
var I *redis.Client // instance

func (o *msRedis) Init(conf *msa.MsConf) {
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
	log.Println("Instantiation for memory storage using redis, status: DONE!!")
}
