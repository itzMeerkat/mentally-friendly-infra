package log

import "go.uber.org/zap"

type Logger interface {
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})
	Infow(string, ...interface{})
	Warnw(string, ...interface{})
	Errorw(string, ...interface{})
	Debugw(string, ...interface{})
}

var GlobalLogger Logger

func InitZapSugared(global bool, production bool, skip int) *zap.SugaredLogger {
	var p *zap.Logger
	var err error
	if production == true {
		p, err = zap.NewProduction(zap.AddCallerSkip(skip))
	} else {
		p, err = zap.NewDevelopment(zap.AddCallerSkip(skip))
	}
	if err != nil {
		panic(err)
	}

	if global {
		GlobalLogger = p.Sugar()
		return nil
	} else {
		return p.Sugar()
	}
}

func Infof(t string, args ...interface{}) {
	GlobalLogger.Infof(t, args)
}
func Warnf(t string, args ...interface{}) {
	GlobalLogger.Warnf(t, args)
}
func Errorf(t string, args ...interface{}) {
	GlobalLogger.Errorf(t, args)
}
func Debugf(t string, args ...interface{}) {
	GlobalLogger.Debugf(t, args)
}

func Infow(msg string, args ...interface{}) {
	GlobalLogger.Infow(msg, args...)
}
func Warnw(msg string, args ...interface{}) {
	GlobalLogger.Warnw(msg, args...)
}
func Errorw(msg string, args ...interface{}) {
	GlobalLogger.Errorw(msg, args...)
}
func Debugw(msg string, args ...interface{}) {
	GlobalLogger.Debugw(msg, args...)
}