package logging

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"io"
	"log"
	"os"
	"strings"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	ProgramName = strings.Split(os.Args[0], "/")[len(strings.Split(os.Args[0], "/"))-1]
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func init() {
	fileLogger := &lumberjack.Logger{
		Filename:   "/tmp/test.log",
		MaxSize:    500, // megabytes
		MaxAge:     28,  // days
		MaxBackups: 3,
	}
	multiWriter := io.MultiWriter(os.Stdout, fileLogger)
	log.SetOutput(multiWriter)
	InfoLogger = log.New(multiWriter, "", 0)
	InfoLogger = log.New(multiWriter, "", 0)
}

func InitLoggers(logFilePath string) {
	fileLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    500, // megabytes
		MaxAge:     28,  // days
		MaxBackups: 3,
	}
	multiWriter := io.MultiWriter(os.Stdout, fileLogger)
	log.SetOutput(multiWriter)
	InfoLogger = log.New(multiWriter, "", 0)
	ErrorLogger = log.New(multiWriter, "", 0)
}

func customLogFormat(logger *log.Logger, level string, colorLevel string, v ...interface{}) {
	logger.SetFlags(log.Ldate | log.Ltime)
	message := fmt.Sprintf("%s%s%s %s%s%s - %s%s%s", colorBlue, ProgramName, colorReset, colorLevel, level, colorReset, colorYellow, fmt.Sprint(v...), colorReset)
	err := logger.Output(2, message)
	if err != nil {
		return
	}
}

func Info(v ...interface{}) {
	customLogFormat(InfoLogger, "INFO", colorGreen, v...)
}

func Error(v ...interface{}) {
	customLogFormat(ErrorLogger, "ERROR", colorRed, v...)
}
