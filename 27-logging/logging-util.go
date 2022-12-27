package main

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Loglevel uint16

const (
	//定义日志级别
	DEBUG Loglevel = iota
	TRACE
	ERROR
	WARN
	INFO
	UNKNOWN
)

func parseLogLevel(s string) (Loglevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "error":
		return ERROR, nil
	case "warn":
		return WARN, nil
	case "info":
		return INFO, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

type logger struct {
	Level Loglevel
}

func NewLog(levelStr string) logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return logger{
		Level: level,
	}
}

func (l logger) enable(logLevel Loglevel) bool {
	return logLevel >= l.Level
}
func (l logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [Error] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}

}
func (l logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [Debug]%s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l logger) Warning(msg string) {
	if l.enable(WARN) {
		now := time.Now()
		fmt.Printf("[%s] [Warnning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [Info]%s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func getInfo(n int) {
	_, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	fmt.Println(file)
	fmt.Println(line)
}
func main() {
	log := NewLog("error")
	getInfo(0)
	for {
		time.Sleep(1 * time.Second)
		log.Debug("This is a debug log")
		log.Warning("This is a warning log")
		log.Error("This is a error log")
		log.Info("This is a info log")
	}
}
