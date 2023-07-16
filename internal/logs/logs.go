package logs

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const myTimeFormat = "02/01/2006 15:04:05"

func WriteLogFile(logType string, information string) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.TraceLevel
	logFile, _ := os.OpenFile("./logg.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.Out = logFile
	if logType == "INFO" {
		log.WithFields(logrus.Fields{
			"env":   "dev",
			"layer": "log",
			"type":  logType,
		}).Info(information)
		logApplicationStdout(logType, information)
	}
	if logType == "ERROR" {
		log.WithFields(logrus.Fields{
			"env":   "dev",
			"layer": "log",
			"type":  logType,
		}).Error(information)
		logApplicationStdout(logType, information)
	}
	if logType == "WARNING" {
		log.WithFields(logrus.Fields{
			"env":   "dev",
			"layer": "log",
			"type":  logType,
		}).Warning(information)
		logApplicationStdout(logType, information)
	}
	defer logFile.Close()
}

func logApplicationStdout(logType string, information string) {
	logString := fmt.Sprintf("[%s]: %s\n", logType, information)
	fmt.Print(logString)
}
