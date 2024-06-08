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

func (o *loggerZerolog) Init(conf *l.LoggerConf, app *a.AppConf) {
	o.level = l.Level(conf.Level)

	ctx := log.With()
	if !conf.HideTime {
		ctx = ctx.Timestamp()
		if conf.FormatTime {
			zerolog.TimeFieldFormat = time.RFC3339
		} else {
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		}
	}
	if conf.HideLevel {
		zerolog.LevelFieldName = ""
	}
	log.Logger = ctx.Logger()

	o.level = l.Level(conf.Level)
	if o.level == l.LInfo {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if o.level == l.LWarning {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if o.level == l.LError {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else if o.level == l.LPanic {
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	} else if o.level == l.LFatal {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

}

// Debug() LoggerItf
func (o *loggerZerolog) Debug() l.LoggerItf {
	o.l = log.Debug()
	return o
}

// Info() LoggerItf
func (o *loggerZerolog) Info() l.LoggerItf {
	o.l = log.Info()
	return o
}

// Warning() LoggerItf
func (o *loggerZerolog) Warning() l.LoggerItf {
	o.l = log.Warn()
	return o
}

// Error() LoggerItf
func (o *loggerZerolog) Error() l.LoggerItf {
	o.l = log.Error()
	return o
}

// Panic() LoggerItf
func (o *loggerZerolog) Panic() l.LoggerItf {
	o.l = log.Debug()
	return o
}

// Fatal() LoggerItf
func (o *loggerZerolog) Fatal() l.LoggerItf {
	o.l = log.Debug()
	return o
}

// Bool(string, bool) LoggerItf
func (o *loggerZerolog) Bool(key string, val bool) l.LoggerItf {
	// if !o.set {
	// 	o = &loggerZerolog{set: true, level: o.level, l: log.Debug()}
	// }
	o.l = o.l.Bool(key, val)
	return o
}

// Int(string, int) LoggerItf
func (o *loggerZerolog) Int(key string, val int) l.LoggerItf {
	o.l = o.l.Int(key, val)
	return o
}

// Int8(string, int8) LoggerItf
func (o *loggerZerolog) Int8(key string, val int8) l.LoggerItf {
	o.l = o.l.Int8(key, val)
	return o
}

// Int16(string, int16) LoggerItf
func (o *loggerZerolog) Int16(key string, val int16) l.LoggerItf {
	o.l = o.l.Int16(key, val)
	return o
}

// Int32(string, int32) LoggerItf
func (o *loggerZerolog) Int32(key string, val int32) l.LoggerItf {
	o.l = o.l.Int32(key, val)
	return o
}

// Int64(string, int64) LoggerItf
func (o *loggerZerolog) Int64(key string, val int64) l.LoggerItf {
	o.l = o.l.Int64(key, val)
	return o
}

// Uint(string, uint) LoggerItf
func (o *loggerZerolog) Uint(key string, val uint) l.LoggerItf {
	o.l = o.l.Uint(key, val)
	return o
}

// Uint8(string, uint8) LoggerItf
func (o *loggerZerolog) Uint8(key string, val uint8) l.LoggerItf {
	o.l = o.l.Uint8(key, val)
	return o
}

// Uint16(string, uint16) LoggerItf
func (o *loggerZerolog) Uint16(key string, val uint16) l.LoggerItf {
	o.l = o.l.Uint16(key, val)
	return o
}

// Uint32(string, uint32) LoggerItf
func (o *loggerZerolog) Uint32(key string, val uint32) l.LoggerItf {
	o.l = o.l.Uint32(key, val)
	return o
}

// Uint64(string, uint64) LoggerItf
func (o *loggerZerolog) Uint64(key string, val uint64) l.LoggerItf {
	o.l = o.l.Uint64(key, val)
	return o
}

// String(string, string) LoggerItf
func (o *loggerZerolog) String(key string, val string) l.LoggerItf {
	o.l = o.l.Str(key, val)
	return o
}

// Send()
func (o *loggerZerolog) Send() {
	o.l.Send()
}
