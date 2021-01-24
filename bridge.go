package logrusmiddleware

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Logger bridges the echo.Logger interface to use sirupsen.Logrus as the implementation
type Logger struct {
	*logrus.Logger
}

//Output gets the output
func (l Logger) Output() io.Writer {
	return l.Logger.Out
}

//SetOutput sets the output
func (l Logger) SetOutput(w io.Writer) {
	l.Logger.Out = w
}

// Prefix is not supported and returns an empty string
func (l Logger) Prefix() string {
	return ""
}

// SetPrefix is not supported and is a NOOP
func (l Logger) SetPrefix(s string) {}

// Level gets the current logging level
func (l Logger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Logger.Panic("Invalid level")
	}

	return log.OFF
}

// SetLevel sets the logging level
func (l Logger) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	default:
		l.Logger.Panic("Invalid level")
	}
}

// SetHeader is not supported and is a NOOP
func (l Logger) SetHeader(header string) {}

//Printj prints with fields
func (l Logger) Printj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Print()
}

//Debugj logs with fields at debug
func (l Logger) Debugj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Debug()
}

//Infoj logs with fields at info
func (l Logger) Infoj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Info()
}

//Warnj logs with fields at warn
func (l Logger) Warnj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Warn()
}

//Errorj logs with fields at error
func (l Logger) Errorj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Error()
}

//Fatalj logs with fields at fatal
func (l Logger) Fatalj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Fatal()
}

//Panicj logs with fields at panic
func (l Logger) Panicj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Panic()
}
