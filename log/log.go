package log

import "go.uber.org/zap"

type Logger interface {
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})
}

var IsGlobal bool
var globalLogger Logger

func InitZapSugared(global bool, production bool) (*zap.SugaredLogger, error) {
	var p *zap.Logger
	var err error
	if production == true {
		p, err = zap.NewProduction(zap.AddCallerSkip(1))
	} else {
		p, err = zap.NewDevelopment(zap.AddCallerSkip(1))
	}
	if err != nil {
		return nil, err
	}

	if global {
		IsGlobal = true
		globalLogger = p.Sugar()
		return nil, nil
	} else {
		IsGlobal = false
		return p.Sugar(), nil
	}
}

func Infof(t string, args ...interface{}) {
	globalLogger.Infof(t, args)
}
func Warnf(t string, args ...interface{}) {
	globalLogger.Warnf(t, args)
}
func Errorf(t string, args ...interface{}) {
	globalLogger.Errorf(t, args)
}
func Debugf(t string, args ...interface{}) {
	globalLogger.Debugf(t, args)
}