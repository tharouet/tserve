package logger

import (
	"github.com/sirupsen/logrus"
)

var DebugMode bool

type Logger interface {
	New() (l Logit)
}

type Logit struct {
	Log *logrus.Logger
}

func New() (l Logit) {
	l.Log = logrus.New()
	if DebugMode {
		l.Log.SetLevel(logrus.DebugLevel)
	}
	return
}

func NewWithTag(tag string) (l Logit) {
	l.Log = logrus.New()
	l.Log.WithField("@tag", tag)
	return
}
