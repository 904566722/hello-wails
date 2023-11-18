package log

import "github.com/sirupsen/logrus"

var Log *logrus.Logger

func Init() error {
	Log = logrus.New()
	Log.SetLevel(logrus.DebugLevel)
	return nil
}
