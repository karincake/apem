package loggerzap

import (
	"go.uber.org/zap"
)

type LoggerConf struct {
	Mode  string
	Level int8
}

var I *zap.Logger // instance

func Init(conf LoggerConf) {
	var err error
	I = &zap.Logger{}
	if conf.Mode == "development" {
		I, err = zap.NewDevelopment()
	} else {
		I, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	defer I.Sync()
	I.Info("instantiation", zap.String("feature", "logger"), zap.String("source", "zap"), zap.String("status", "done"))
}
