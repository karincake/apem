package msredis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"

	lz "github.com/karincake/apem/loggerzap"
)

type MsConf struct {
	Dsn string
}

var I *redis.Client // instance

func Init(conf MsConf) {
	if len(conf.Dsn) == 0 {
		lz.I.Warn("instantiation", zap.String("feature", "memstorage"), zap.String("source", "redis"), zap.String("status", "skipped"))
		return
	}

	I = redis.NewClient(&redis.Options{
		Addr: conf.Dsn,
	})
	_, err := I.Ping().Result()
	if err != nil {
		panic(err)
	}
	lz.I.Info("instantiation", zap.String("feature", "memstorage"), zap.String("source", "redis"), zap.String("status", "done"))
}
