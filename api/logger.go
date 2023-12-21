package api

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	file, err := os.OpenFile(os.Getenv("LOG_FILE_PATH"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	formatter := &logrus.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
		DisableColors:   false,
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			function := ""
			file := fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
			return function, file
		},
	}

	log := logrus.New()
	log.SetFormatter(formatter)
	log.SetReportCaller(true)
	log.SetOutput(file)

	return log
}
