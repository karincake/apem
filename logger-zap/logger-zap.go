package loggerzap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	a "github.com/karincake/apem/appa"
	l "github.com/karincake/apem/loggera"
)

type loggerZap struct {
	set       bool
	level     l.Level
	execLevel l.Level
	fields    []zap.Field
}

var O loggerZap
var I *zap.Logger

// Init(*LoggerCfg, *appa.AppCfg)
func (obj *loggerZap) Init(conf *l.LoggerCfg, app *a.AppCfg) {
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.MessageKey = zapcore.OmitKey
	if conf.HideLevel {
		encoderCfg.LevelKey = zapcore.OmitKey
	}
	if conf.FormatTime {
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	I = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer I.Sync()

	// different leveling, can't do autocast
	obj.level = l.Level(conf.Level)
	if obj.level == l.LInfo {
		atom.SetLevel(zapcore.InfoLevel)
	} else if obj.level == l.LWarning {
		atom.SetLevel(zapcore.WarnLevel)
	} else if obj.level == l.LError {
		atom.SetLevel(zapcore.ErrorLevel)
	} else if obj.level == l.LPanic {
		atom.SetLevel(zapcore.PanicLevel)
	} else if obj.level == l.LFatal {
		atom.SetLevel(zapcore.FatalLevel)
	} else {
		atom.SetLevel(zapcore.DebugLevel)
	}

	obj.Info().String("scope", "instantiation").String("source", "zap").String("status", "created").Send()
}

// Debug() LoggerItf
func (obj *loggerZap) Debug() l.LoggerItf {
	return &loggerZap{set: true, level: obj.level, execLevel: l.LDebug}
}

// Info() LoggerItf
func (obj *loggerZap) Info() l.LoggerItf {
	return &loggerZap{set: true, level: obj.level, execLevel: l.LInfo}
}

// Warning() LoggerItf
func (obj *loggerZap) Warning() l.LoggerItf {
	return &loggerZap{set: true, level: obj.level, execLevel: l.LWarning}
}

// Error() LoggerItf
func (obj *loggerZap) Error() l.LoggerItf {
	return &loggerZap{set: true, level: obj.level, execLevel: l.LError}
}

// Panic() LoggerItf
func (obj *loggerZap) Panic() l.LoggerItf {
	return &loggerZap{set: true, level: obj.level, execLevel: l.LPanic}
}

// Fatal() LoggerItf
func (obj *loggerZap) Fatal() l.LoggerItf {
	return &loggerZap{set: true, level: obj.level, execLevel: l.LFatal}
}

// Bool(string, bool) LoggerItf
func (obj *loggerZap) Bool(key string, val bool) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Bool(key, val))
	return obj
}

// Int(string, int) LoggerItf
func (obj *loggerZap) Int(key string, val int) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Int(key, val))
	return obj
}

// Int16(string, int16) LoggerItf
func (obj *loggerZap) Int8(key string, val int8) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Int8(key, val))
	return obj
}

// Int8(string, int8) LoggerItf
func (obj *loggerZap) Int16(key string, val int16) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Int16(key, val))
	return obj
}

// Int32(string, int32) LoggerItf
func (obj *loggerZap) Int32(key string, val int32) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Int32(key, val))
	return obj
}

// Int64(string, int64) LoggerItf
func (obj *loggerZap) Int64(key string, val int64) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Int64(key, val))
	return obj
}

// Uint(string, uint) LoggerItf
func (obj *loggerZap) Uint(key string, val uint) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Uint(key, val))
	return obj
}

// Uint8(string, uint8) LoggerItf
func (obj *loggerZap) Uint8(key string, val uint8) l.LoggerItf {
	obj.fields = append(obj.fields, zap.Uint8(key, val))
	return obj
}

// Uint16(string, uint16) LoggerItf
func (obj *loggerZap) Uint16(key string, val uint16) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Uint16(key, val))
	return obj
}

// Uint32(string, uint32) LoggerItf
func (obj *loggerZap) Uint32(key string, val uint32) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Uint32(key, val))
	return obj
}

// Uint64(string, uint64) LoggerItf
func (obj *loggerZap) Uint64(key string, val uint64) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.Uint64(key, val))
	return obj
}

// String(string, string) LoggerItf
func (obj *loggerZap) String(key string, val string) l.LoggerItf {
	if !obj.set {
		obj = &loggerZap{set: true, level: obj.level}
	}
	obj.fields = append(obj.fields, zap.String(key, val))
	return obj
}

// Send()
func (obj *loggerZap) Send() {
	if obj.execLevel == l.LDebug {
		I.Debug("", obj.fields...)
	} else if obj.execLevel == l.LInfo {
		I.Info("", obj.fields...)
	} else if obj.execLevel == l.LWarning {
		I.Warn("", obj.fields...)
	} else if obj.execLevel == l.LError {
		I.Error("", obj.fields...)
	} else if obj.execLevel == l.LPanic {
		I.Panic("", obj.fields...)
	} else if obj.execLevel == l.LFatal {
		I.Fatal("", obj.fields...)
	}
}
