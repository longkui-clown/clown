package core

import "fmt"

type Logger interface {
	DEBUG(format string, args ...any)
	INFO(format string, args ...any)
	WARNING(format string, args ...any)
	ERROR(format string, args ...any)
}

var _ Logger = (*DefaultManagerLogger)(nil)

type DefaultManagerLogger struct{}

func (l *DefaultManagerLogger) DEBUG(format string, args ...any) {
	l.log("Debug", format, args...)
}

func (l *DefaultManagerLogger) INFO(format string, args ...any) {
	l.log("Info", format, args...)
}

func (l *DefaultManagerLogger) WARNING(format string, args ...any) {
	l.log("Warning", format, args...)
}

func (l *DefaultManagerLogger) ERROR(format string, args ...any) {
	l.log("Error", format, args...)
}

func (l *DefaultManagerLogger) log(tag string, format string, args ...any) {
	fmt.Printf("["+tag+"] "+format+"\n", args...)
}
