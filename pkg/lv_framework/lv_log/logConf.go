package lv_log

import (
	"fmt"
	"io"
)

type ILog interface {
	Error(args ...any)
	ErrorTraceId(traceId any, args ...any)
	Fatal(args ...any)
	FatalTraceId(traceId any, args ...any)
	Warn(args ...any)
	WarnTraceId(traceId any, args ...any)
	Info(args ...any)
	InfoTraceId(traceId any, args ...any)
	Debug(args ...any)
	DebugTraceId(traceId any, args ...any)
	Errorf(format string, args ...any)
	Warnf(format string, args ...any)
	Infof(format string, args ...any)
	Debugf(format string, args ...any)
	GetLogWriter() io.Writer
}

var Log ILog //主数据库

func GetLog() ILog {
	if Log == nil {
		fmt.Println("log is nil!!!!!!!")
	}
	return Log
}

func Error(args ...any) {
	if Log != nil {
		Log.Error(args)
	} else {
		fmt.Println(args)
	}
}

func ErrorTraceId(traceId any, args ...any) {
	if Log != nil {
		Log.Error(args)
	} else {
		fmt.Println(args)
	}
}
func Fatal(args ...any) {
	Log.Fatal(args)
}
func FatalTraceId(traceId any, args ...any) {
	if Log != nil {
		Log.Fatal(args)
	} else {
		fmt.Println(args)
	}
}
func Warn(args ...any) {
	if Log != nil {
		Log.Warn(args)
	} else {
		fmt.Println(args)
	}
}
func WarnTraceId(traceId any, args ...any) {
	if Log != nil {
		Log.Warn(args)
	} else {
		fmt.Println(args)
	}
}
func Info(args ...any) {
	if Log != nil {
		Log.Info(args)
	} else {
		fmt.Println(args)
	}
}
func InfoTraceId(traceId any, args ...any) {
	if Log != nil {
		Log.Info(args)
	} else {
		fmt.Println(args)
	}
}
func Debug(args ...any) {
	if Log != nil {
		Log.Debug(args)
	} else {
		fmt.Println(args)
	}
}
func DebugTraceId(traceId any, args ...any) {
	if Log != nil {
		Log.Debug(args)
	} else {
		fmt.Println(args)
	}
}
func Errorf(format string, args ...any) {
	if Log != nil {
		Log.Errorf(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}
func Warnf(format string, args ...any) {
	if Log != nil {
		Log.Warnf(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}
func Infof(format string, args ...any) {
	if Log != nil {
		Log.Infof(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}

func Debugf(format string, args ...any) {
	if Log != nil {
		Log.Debugf(format, args)
	} else {
		fmt.Printf("\n"+format, args)
	}
}
