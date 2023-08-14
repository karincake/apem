package loggerzap

import (
	"go.uber.org/zap"
)

// Configuration type that is used by the core
type LoggerConf struct {
	Mode  string
	Level int8
}

// Instance of the logger
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
