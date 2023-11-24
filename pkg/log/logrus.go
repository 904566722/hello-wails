package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"changeme/pkg/consts"
)

var logFile = filepath.Join(consts.AppFilePath, consts.LogName)

var Log *logrus.Logger

func Init() error {
	Log = logrus.New()
	Log.SetLevel(logrus.DebugLevel)
	Log.SetReportCaller(false)

	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		Log.Errorf("open log file failed: %v", err)
		return err
	}
	Log.SetOutput(io.MultiWriter(file, os.Stdout))
	Log.Infof("log file: %s", logFile)

	return nil
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	Log.Infoln(args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	Log.Debugln(args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	Log.Warnln(args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	Log.Errorln(args...)
}

func InfoWithFields(fields map[string]interface{}, args ...interface{}) {
	Log.WithFields(fields).Info(args...)
}

func DebugWithFields(fields map[string]interface{}, args ...interface{}) {
	Log.WithFields(fields).Debug(args...)
}

func WarnWithFields(fields map[string]interface{}, args ...interface{}) {
	Log.WithFields(fields).Warn(args...)
}

func ErrorWithFields(fields map[string]interface{}, args ...interface{}) {
	Log.WithFields(fields).Error(args...)
}
