package xxljob

import (
	"log"
)

type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	log.Printf("自定义日志 - "+format, a...)
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Printf("自定义日志 - "+format, a...)
}
