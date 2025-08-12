package loggerzerolog

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	a "github.com/karincake/apem/appa"
	l "github.com/karincake/apem/loggera"
)

type loggerZerolog struct {
	// set       bool
	level l.Level
	// execLevel l.Level
	l *zerolog.Event
}

type KeyValFunc func(key string, val any)

var O loggerZerolog
var Ctx zerolog.Context

func (obj *loggerZerolog) Init(conf *l.LoggerCfg, app *a.AppCfg) {
	obj.level = l.Level(conf.Level)

	if !conf.HideTime {
		if conf.FormatTime {
			zerolog.TimeFieldFormat = time.RFC3339
		} else {
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		}
	}
	if conf.HideLevel {
		zerolog.LevelFieldName = ""
	}
	log.Logger = Ctx.Logger()

	obj.level = l.Level(conf.Level)
	if obj.level == l.LInfo {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if obj.level == l.LWarning {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if obj.level == l.LError {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else if obj.level == l.LPanic {
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	} else if obj.level == l.LFatal {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

}

// Debug() LoggerItf
func (obj *loggerZerolog) Debug() l.LoggerItf {
	obj.l = log.Debug()
	return obj
}

// Info() LoggerItf
func (obj *loggerZerolog) Info() l.LoggerItf {
	obj.l = log.Info()
	return obj
}

// Warning() LoggerItf
func (obj *loggerZerolog) Warning() l.LoggerItf {
	obj.l = log.Warn()
	return obj
}

// Error() LoggerItf
func (obj *loggerZerolog) Error() l.LoggerItf {
	obj.l = log.Error()
	return obj
}

// Panic() LoggerItf
func (obj *loggerZerolog) Panic() l.LoggerItf {
	obj.l = log.Debug()
	return obj
}

// Fatal() LoggerItf
func (obj *loggerZerolog) Fatal() l.LoggerItf {
	obj.l = log.Debug()
	return obj
}

// Bool(string, bool) LoggerItf
func (obj *loggerZerolog) Bool(key string, val bool) l.LoggerItf {
	// if !o.set {
	// 	o = &loggerZerolog{set: true, level: obj.level, l: log.Debug()}
	// }
	obj.l = obj.l.Bool(key, val)
	return obj
}

// Int(string, int) LoggerItf
func (obj *loggerZerolog) Int(key string, val int) l.LoggerItf {
	obj.l = obj.l.Int(key, val)
	return obj
}

// Int8(string, int8) LoggerItf
func (obj *loggerZerolog) Int8(key string, val int8) l.LoggerItf {
	obj.l = obj.l.Int8(key, val)
	return obj
}

// Int16(string, int16) LoggerItf
func (obj *loggerZerolog) Int16(key string, val int16) l.LoggerItf {
	obj.l = obj.l.Int16(key, val)
	return obj
}

// Int32(string, int32) LoggerItf
func (obj *loggerZerolog) Int32(key string, val int32) l.LoggerItf {
	obj.l = obj.l.Int32(key, val)
	return obj
}

// Int64(string, int64) LoggerItf
func (obj *loggerZerolog) Int64(key string, val int64) l.LoggerItf {
	obj.l = obj.l.Int64(key, val)
	return obj
}

// Uint(string, uint) LoggerItf
func (obj *loggerZerolog) Uint(key string, val uint) l.LoggerItf {
	obj.l = obj.l.Uint(key, val)
	return obj
}

// Uint8(string, uint8) LoggerItf
func (obj *loggerZerolog) Uint8(key string, val uint8) l.LoggerItf {
	obj.l = obj.l.Uint8(key, val)
	return obj
}

// Uint16(string, uint16) LoggerItf
func (obj *loggerZerolog) Uint16(key string, val uint16) l.LoggerItf {
	obj.l = obj.l.Uint16(key, val)
	return obj
}

// Uint32(string, uint32) LoggerItf
func (obj *loggerZerolog) Uint32(key string, val uint32) l.LoggerItf {
	obj.l = obj.l.Uint32(key, val)
	return obj
}

// Uint64(string, uint64) LoggerItf
func (obj *loggerZerolog) Uint64(key string, val uint64) l.LoggerItf {
	obj.l = obj.l.Uint64(key, val)
	return obj
}

// String(string, string) LoggerItf
func (obj *loggerZerolog) String(key string, val string) l.LoggerItf {
	obj.l = obj.l.Str(key, val)
	return obj
}

// Send()
func (obj *loggerZerolog) Send() {
	obj.l.Send()
}

// Any(string, any) LoggerItf
func (obj *loggerZerolog) Any(key string, val any) l.LoggerItf {
	obj.l = obj.l.Any(key, val)
	return obj
}
