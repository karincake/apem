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

// Init(*LoggerConf, *appa.AppConf)
func (o *loggerZap) Init(conf *l.LoggerConf, app *a.AppConf) {
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
	o.level = l.Level(conf.Level)
	if o.level == l.LInfo {
		atom.SetLevel(zapcore.InfoLevel)
	} else if o.level == l.LWarning {
		atom.SetLevel(zapcore.WarnLevel)
	} else if o.level == l.LError {
		atom.SetLevel(zapcore.ErrorLevel)
	} else if o.level == l.LPanic {
		atom.SetLevel(zapcore.PanicLevel)
	} else if o.level == l.LFatal {
		atom.SetLevel(zapcore.FatalLevel)
	} else {
		atom.SetLevel(zapcore.DebugLevel)
	}

	O.Info().String("scope", "instantiation").String("source", "zap").String("status", "created").Send()
}

// Debug() LoggerItf
func (o *loggerZap) Debug() l.LoggerItf {
	return &loggerZap{set: true, level: o.level, execLevel: l.LDebug}
}

// Info() LoggerItf
func (o *loggerZap) Info() l.LoggerItf {
	return &loggerZap{set: true, level: o.level, execLevel: l.LInfo}
}

// Warning() LoggerItf
func (o *loggerZap) Warning() l.LoggerItf {
	return &loggerZap{set: true, level: o.level, execLevel: l.LWarning}
}

// Error() LoggerItf
func (o *loggerZap) Error() l.LoggerItf {
	return &loggerZap{set: true, level: o.level, execLevel: l.LError}
}

// Panic() LoggerItf
func (o *loggerZap) Panic() l.LoggerItf {
	return &loggerZap{set: true, level: o.level, execLevel: l.LPanic}
}

// Fatal() LoggerItf
func (o *loggerZap) Fatal() l.LoggerItf {
	return &loggerZap{set: true, level: o.level, execLevel: l.LFatal}
}

// Bool(string, bool) LoggerItf
func (o *loggerZap) Bool(key string, val bool) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Bool(key, val))
	return o
}

// Int(string, int) LoggerItf
func (o *loggerZap) Int(key string, val int) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Int(key, val))
	return o
}

// Int16(string, int16) LoggerItf
func (o *loggerZap) Int8(key string, val int8) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Int8(key, val))
	return o
}

// Int8(string, int8) LoggerItf
func (o *loggerZap) Int16(key string, val int16) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Int16(key, val))
	return o
}

// Int32(string, int32) LoggerItf
func (o *loggerZap) Int32(key string, val int32) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Int32(key, val))
	return o
}

// Int64(string, int64) LoggerItf
func (o *loggerZap) Int64(key string, val int64) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Int64(key, val))
	return o
}

// Uint(string, uint) LoggerItf
func (o *loggerZap) Uint(key string, val uint) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Uint(key, val))
	return o
}

// Uint8(string, uint8) LoggerItf
func (o *loggerZap) Uint8(key string, val uint8) l.LoggerItf {
	o.fields = append(o.fields, zap.Uint8(key, val))
	return o
}

// Uint16(string, uint16) LoggerItf
func (o *loggerZap) Uint16(key string, val uint16) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Uint16(key, val))
	return o
}

// Uint32(string, uint32) LoggerItf
func (o *loggerZap) Uint32(key string, val uint32) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Uint32(key, val))
	return o
}

// Uint64(string, uint64) LoggerItf
func (o *loggerZap) Uint64(key string, val uint64) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.Uint64(key, val))
	return o
}

// String(string, string) LoggerItf
func (o *loggerZap) String(key string, val string) l.LoggerItf {
	if !o.set {
		o = &loggerZap{set: true, level: o.level}
	}
	o.fields = append(o.fields, zap.String(key, val))
	return o
}

// Send()
func (o *loggerZap) Send() {
	if o.execLevel == l.LDebug {
		I.Debug("", o.fields...)
	} else if o.execLevel == l.LInfo {
		I.Info("", o.fields...)
	} else if o.execLevel == l.LWarning {
		I.Warn("", o.fields...)
	} else if o.execLevel == l.LError {
		I.Error("", o.fields...)
	} else if o.execLevel == l.LPanic {
		I.Panic("", o.fields...)
	} else if o.execLevel == l.LFatal {
		I.Fatal("", o.fields...)
	}
}
