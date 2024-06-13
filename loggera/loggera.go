package loggera

import (
	"github.com/karincake/apem/appa"
)

type Level int8

type LoggerCfg struct {
	Mode       string
	Level      int8
	HideLevel  bool `yaml:"hideLevel"`
	HideTime   bool `yaml:"hideTime"`
	FormatTime bool `yaml:"formatTime"`
}

type LoggerItf interface {
	Init(*LoggerCfg, *appa.AppCfg)
	Debug() LoggerItf
	Info() LoggerItf
	Warning() LoggerItf
	Error() LoggerItf
	Panic() LoggerItf
	Fatal() LoggerItf
	Bool(string, bool) LoggerItf
	Int(string, int) LoggerItf
	Int8(string, int8) LoggerItf
	Int16(string, int16) LoggerItf
	Int32(string, int32) LoggerItf
	Int64(string, int64) LoggerItf
	Uint(string, uint) LoggerItf
	Uint8(string, uint8) LoggerItf
	Uint16(string, uint16) LoggerItf
	Uint32(string, uint32) LoggerItf
	Uint64(string, uint64) LoggerItf
	String(string, string) LoggerItf
	Send()
}

const (
	LDebug Level = iota
	LInfo
	LWarning
	LError
	LPanic
	LFatal
)

// Bool
// Int
// Uint
// Uint8
// Uint16
// Uint32
// Uint64
// Uintptr
// Float32
// Float64
// Complex64
// Complex128
// Array
// Chan
// Func
// Interface
// Map
// Pointer
// Slice
// String
// Struct
// UnsafePointer
