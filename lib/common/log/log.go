package log

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

var (
	infolnLogger       *log.Logger = nil
	warningLogger      *log.Logger = log.New(color.Error, color.YellowString("WARNING: "), 0)
	errorLogger        *log.Logger = log.New(color.Error, color.RedString("ERROR: "), 0)
	PanicInsteadOfExit bool
)

func Info(format string, v ...any) {
	if infolnLogger != nil {
		infolnLogger.Printf(format, v...)
	}
}

func Printlnf(format string, v ...any) {
	fmt.Printf(format, v...)
	fmt.Println()
}

func Warning(format string, v ...any) {
	warningLogger.Printf(format, v...)
}

func Error(format string, v ...any) {
	errorLogger.Printf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	if PanicInsteadOfExit {
		Error(format, v...)
		panic(fmt.Sprintf(format, v...))
	}
	errorLogger.Fatalf(format, v...)
}
