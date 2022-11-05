package logging

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger    *log.Logger
	logPrefix = ""
)

// Setup initialize the log instance
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix("DEBUG")
	logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix("INFO")
	logger.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix("WARNING")
	logger.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix("ERROR")
	logger.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix("FATAL")
	logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level string) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", level, filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", level)
	}

	logger.SetPrefix(logPrefix)
}
