package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout

	Log.Formatter = &logrus.JSONFormatter{}
}
