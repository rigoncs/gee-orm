package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	errorlog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infolog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorlog, infolog}
	mu       sync.Mutex
)

// log methods
var (
	Error  = errorlog.Println
	Errorf = errorlog.Printf
	Info   = infolog.Println
	Infof  = infolog.Printf
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorlog.SetOutput(io.Discard)
	}
	if InfoLevel < level {
		infolog.SetOutput(io.Discard)
	}
}
