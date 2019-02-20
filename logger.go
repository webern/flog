// iqBid Pricing System, Copyright (c) 2019 by Inteliquent, Inc.

// flog package: use this package instead of logrus or the build in log package. the name is a bit of a joke. we need a
// name that is different from 'log' and 'logrus', and we need something short and easy to type. excessive logging does
// feel a bit like flogging, so we will use the package name 'flog'
package flog

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger logrus.Logger

func New() *Logger {
	logrusLogger := logrus.New()
	return (*Logger)(logrusLogger)
}

var (
	// std is the name of the standard logger in stdlib `log`
	std = New()
)

func StandardLogger() *Logger {
	return (*Logger)(std)
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include bool) {
	// logrus.SetReportCaller(include)
}

// SetLevel sets the standard logger level.
func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

// GetLevel returns the standard logger level.
func GetLevel() logrus.Level {
	return logrus.GetLevel()
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level logrus.Level) bool {
	return logrus.IsLevelEnabled(level)
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	return logrus.WithField(logrus.ErrorKey, err)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return logrus.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.WithFields(fields)
}

// WithTime creats an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *logrus.Entry {
	return logrus.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	if logrus.GetLevel() >= logrus.TraceLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Trace(args...)
	}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if logrus.GetLevel() >= logrus.DebugLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Debug(args...)
	}
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	if logrus.GetLevel() >= logrus.InfoLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Print(args...)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if logrus.GetLevel() >= logrus.InfoLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Info(args...)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if logrus.GetLevel() >= logrus.WarnLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Warn(args...)
	}
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	if logrus.GetLevel() >= logrus.WarnLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Warning(args...)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if logrus.GetLevel() >= logrus.ErrorLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Error(args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if logrus.GetLevel() >= logrus.PanicLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Panic(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	if logrus.GetLevel() >= logrus.FatalLevel {
		caller := BlazeCaller(2) + " - "
		args = append(args, nil)
		moveto := args[1:]
		copy(moveto, args)
		args[0] = caller
		logrus.Fatal(args...)
	}
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.TraceLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Tracef(format, args...)
	}
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.DebugLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Debugf(format, args...)
	}
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.InfoLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Printf(format, args...)
	}
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.InfoLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Infof(format, args...)
	}
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.WarnLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Warnf(format, args...)
	}
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.WarnLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Warningf(format, args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.ErrorLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Errorf(format, args...)
	}
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.PanicLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Panicf(format, args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	if logrus.GetLevel() >= logrus.FatalLevel {
		caller := BlazeCaller(2)
		format = caller + " - " + format
		logrus.Fatalf(format, args...)
	}
}

func InfoAlways(args ...interface{}) {
	theLevel := GetLevel()

	if theLevel != logrus.InfoLevel {
		SetLevel(logrus.InfoLevel)
	}

	caller := BlazeCaller(2) + " - "
	args = append(args, nil)
	moveto := args[1:]
	copy(moveto, args)
	args[0] = caller
	logrus.Info(args...)

	if theLevel != logrus.InfoLevel {
		SetLevel(theLevel)
	}
}

func InfofAlways(format string, args ...interface{}) {
	theLevel := GetLevel()

	if theLevel != logrus.InfoLevel {
		SetLevel(logrus.InfoLevel)
	}

	caller := BlazeCaller(2)
	format = caller + " - " + format
	logrus.Infof(format, args...)

	if theLevel != logrus.InfoLevel {
		SetLevel(theLevel)
	}
}

/// My Own Thing, not Logrus
func LogPlain(msg string, level logrus.Level) {
	switch level {
	case logrus.FatalLevel:
		logrus.Fatal(msg)
	case logrus.PanicLevel:
		logrus.Panic(msg)
	case logrus.TraceLevel:
		logrus.Trace(msg)
	case logrus.DebugLevel:
		logrus.Debug(msg)
	case logrus.ErrorLevel:
		logrus.Error(msg)
	case logrus.WarnLevel:
		logrus.Warn(msg)
	case logrus.InfoLevel:
		logrus.Info(msg)
	default:
		Bug()
	}
}
