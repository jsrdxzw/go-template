package extensions

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

type LogConf struct {
	EnableStdOut bool `json:"enable_std_out"`
	EnableFile   bool `json:"enable_file"`
}

type LoggerExt struct {
	LogConf
	logger logx.Logger
}

func NewLogExt(c LogConf) *LoggerExt {
	return &LoggerExt{LogConf: c}
}

func (l *LoggerExt) Close() error {
	return logx.Close()
}

func (l *LoggerExt) Init() error {
	l.logger = logx.WithContext(context.Background())
	return nil
}

func (l *LoggerExt) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LoggerExt) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LoggerExt) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LoggerExt) Panic(e error) {
	l.logger.Error(e)
	os.Exit(1)
}
