package log

import (
	"fmt"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

// Predefined handlers
var (
	root          *logger
	StdoutHandler = StreamHandler(os.Stdout, LogfmtWithGIDFormat())
	StderrHandler = StreamHandler(os.Stderr, LogfmtWithGIDFormat())
)

func init() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		StdoutHandler = StreamHandler(colorable.NewColorableStdout(), TerminalFormat())
	}

	if isatty.IsTerminal(os.Stderr.Fd()) {
		StderrHandler = StreamHandler(colorable.NewColorableStderr(), TerminalFormat())
	}

	root = &logger{[]interface{}{}, new(swapHandler)}
	root.SetHandler(StdoutHandler)
}

// New returns a new logger with the given context.
// New is a convenient alias for Root().New
func New(ctx ...interface{}) Logger {
	return root.New(ctx...)
}

// Root returns the root logger
func Root() Logger {
	return root
}

// The following functions bypass the exported logger methods (logger.Debug,
// etc.) to keep the call depth the same for all paths to logger.write so
// runtime.Caller(2) always refers to the call site in client code.

// Debug is a convenient alias for Root().Debug
func Debug(msg string, ctx ...interface{}) {
	root.write(msg, LvlDebug, ctx)
}

// Info is a convenient alias for Root().Info
func Info(msg string, ctx ...interface{}) {
	root.write(msg, LvlInfo, ctx)
}

// Warn is a convenient alias for Root().Warn
func Warn(msg string, ctx ...interface{}) {
	root.write(msg, LvlWarn, ctx)
}

// Error is a convenient alias for Root().Error
func Error(msg string, ctx ...interface{}) {
	root.write(msg, LvlError, ctx)
}

// Crit is a convenient alias for Root().Crit
func Crit(msg string, ctx ...interface{}) {
	root.write(msg, LvlCrit, ctx)
}

// Fatal is a convenient alias for Root().Crit
var Fatal = Crit

func Debugf(format string, a ...interface{}) {
	msg :=  fmt.Sprintf(format, a...)
	Debug(msg)
}

func Infof(format string, a ...interface{}) {
	msg :=  fmt.Sprintf(format, a...)
	Info(msg)
}

func Warnf(format string, a ...interface{}) {
	msg :=  fmt.Sprintf(format, a...)
	Warn(msg)
}

func Errorf(format string, a ...interface{}) {
	msg :=  fmt.Sprintf(format, a...)
	Error(msg)
}

func Fatalf(format string, a ...interface{}) {
	msg :=  fmt.Sprintf(format, a...)
	Fatal(msg)
}

func Critf(format string, a ...interface{}) {
	msg :=  fmt.Sprintf(format, a...)
	Crit(msg)
}
