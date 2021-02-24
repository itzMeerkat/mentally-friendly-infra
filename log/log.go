package log

import "go.uber.org/zap"

type Logger interface {
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})
}

var GlobalLogger Logger

func InitZapSugared(global bool, production bool) *zap.SugaredLogger {
	var p *zap.Logger
	var err error
	if production == true {
		p, err = zap.NewProduction(zap.AddCallerSkip(1))
	} else {
		p, err = zap.NewDevelopment(zap.AddCallerSkip(1))
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