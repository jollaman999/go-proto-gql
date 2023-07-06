package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var customLogger *log.Logger
var fpLog *os.File

const (
	// INFO : Informational messages printed with green.
	INFO = "INFO"
	// DEBUG : Debugging messages printed with teal.
	DEBUG = "DEBUG"
	// WARN : Warning messages printed with yellow.
	WARN = "WARN"
	// ERROR : Error messages printed with red.
	ERROR = "ERROR"
	// CRITICAL : Critical messages printed with magenta.
	CRITICAL = "CRITICAL"
	// NONE : Print normal messages with none of color and date, prefix.
	NONE = "NONE"
)

var (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	magenta = "\033[35m"
	teal    = "\033[36m"
)

var date string
var _time string

func getDateAndTime() {
	now := time.Now()

	year := fmt.Sprintf("%04d", now.Year())
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())

	hour := fmt.Sprintf("%02d", now.Hour())
	minute := fmt.Sprintf("%02d", now.Minute())
	second := fmt.Sprintf("%02d", now.Second())

	date = year + "/" + month + "/" + day
	_time = hour + ":" + minute + ":" + second
}

func getPrefix(logLevel string) (prefixForPrint string, prefixForWrite string) {
	getDateAndTime()

	switch logLevel {
	case INFO:
		return date + " " + _time + " [ " + green + "INFO" + reset + " ] ", "[ INFO ] "
	case DEBUG:
		return date + " " + _time + " [ " + teal + "DEBUG" + reset + " ] ", "[ DEBUG ] "
	case WARN:
		return date + " " + _time + " [ " + yellow + "WARN" + reset + " ] ", "[ WARN ] "
	case ERROR:
		return date + " " + _time + " [ " + red + "ERROR" + reset + " ] ", "[ ERROR ] "
	case CRITICAL:
		return date + " " + _time + " [ " + magenta + "CRITICAL" + reset + " ] ", "[ CRITICAL ] "
	default:
		return "", ""
	}
}

func getCallLocation(printCallLocation bool) string {
	if printCallLocation {
		skip := 3
		if strings.HasSuffix(os.Args[0], ".test") {
			skip = 2
		}
		_, filepath, line, _ := runtime.Caller(skip)
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filename := strings.Replace(filepath, pwd+"/", "", -1)
		return filename + ":" + strconv.Itoa(line) + ": "
	}

	return ""
}

// Println : Print the log with a colored level with new line
func Println(logLevel string, printCallLocation bool, msg ...interface{}) {
	prefixForPrint, prefixForWrite := getPrefix(logLevel)
	fmt.Println(prefixForPrint + getCallLocation(printCallLocation) + fmt.Sprint(msg...))
	if customLogger == nil {
		return
	}
	customLogger.Println(prefixForWrite + getCallLocation(printCallLocation) + fmt.Sprint(msg...))
}

// Printf : Print the formatted log with a colored level
func Printf(logLevel string, printCallLocation bool, format string, a ...any) {
	prefixForPrint, prefixForWrite := getPrefix(logLevel)
	fmt.Printf(prefixForPrint+getCallLocation(printCallLocation)+format, a...)
	if customLogger == nil {
		return
	}
	customLogger.Printf(prefixForWrite+getCallLocation(printCallLocation)+format, a...)
}

// GetLogger : Return the custom logger
func GetLogger() *log.Logger {
	return customLogger
}

// SetLogger : Set the custom logger
func SetLogger(logger *log.Logger) {
	customLogger = logger
}
