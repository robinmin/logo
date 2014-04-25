// Package logo provides global logging methods so that multiple packages may
// log to a single application-defined stream.
//
// If any logging methods are called before SetOutput is called, nothing will be outputted.
package logo

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	CRITICAL = 1 << iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
	numLevels int = iota
	ALL       int = CRITICAL | ERROR | WARNING | NOTICE | INFO | DEBUG

	logFlags int = log.Ltime | log.Lshortfile
)

type LevelBasedLogger struct {
	Module    string
	Logger    *log.Logger
	LevelMask int
}

var (
	loggers map[string]*LevelBasedLogger
	mut     *sync.RWMutex

	levelStrings = map[int]string{
		CRITICAL: "C",
		ERROR:    "E",
		WARNING:  "W",
		NOTICE:   "N",
		INFO:     "I",
		DEBUG:    "D",
		ALL:      "A",
	}
)

func init() {
	loggers = make(map[string]*LevelBasedLogger)
	mut = &sync.RWMutex{}

	// add os.stdout as the default output
	AddLogger("stdout", os.Stdout, ALL)
}

func AddLogger(module string, w io.Writer, levelMask int) *log.Logger {
	mut.Lock()
	defer mut.Unlock()

	if _, ok := loggers[module]; !ok {
		// lgr := log.New(w, logFlags, logFlags)
		lgr := log.New(w, "", logFlags)
		loggers[module] = &LevelBasedLogger{
			Module:    module,
			Logger:    lgr,
			LevelMask: levelMask,
		}
		return lgr
	} else {
		return nil
	}
}

func ReleaseLogger(module string) bool {
	mut.Lock()
	defer mut.Unlock()

	if _, ok := loggers[module]; !ok {
		return false
	} else {
		delete(loggers, module)
		return true
	}
}

func write(msg string, level int) {
	for _, logger := range loggers {
		if logger.LevelMask&level == level {
			logger.Logger.Output(3, "["+levelStrings[level]+"] "+msg)
		}
	}
}

func Critical(format string, v ...interface{}) {
	write(fmt.Sprintf(format, v...), CRITICAL)
}

func Error(format string, v ...interface{}) {
	write(fmt.Sprintf(format, v...), ERROR)
}

func Warn(format string, v ...interface{}) {
	write(fmt.Sprintf(format, v...), WARNING)
}

func Notice(format string, v ...interface{}) {
	write(fmt.Sprintf(format, v...), NOTICE)
}

func Info(format string, v ...interface{}) {
	write(fmt.Sprintf(format, v...), INFO)
}

func Debug(format string, v ...interface{}) {
	write(fmt.Sprintf(format, v...), DEBUG)
}
