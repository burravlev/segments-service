/*
* @File: pkg.logger.logger.go
* @Desctiption: setting up application logger
* @Author: Daniil Buravlev (burravlev.d.a@gmail.com)
 */
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init(level string) {
	file, err := os.OpenFile("avito.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %s", err)
	}
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrusLevel)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(file)
}
