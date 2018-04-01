package main

import (
	_ "expvar"
	"net/http"
	_ "sea-robber-api/app/utils/cache"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"sea-robber-api/app/config"
	"sea-robber-api/app/model"
	"sea-robber-api/app/router"
)

func main() {
	go http.ListenAndServe(":8080", nil)
	initLogLevel()
	logrus.Warnf("DB => %s", config.GetString("db.mysql"))
	model.OpenDB(config.GetString("db.mysql"))
	defer model.CloseDB()
	router.Run(8000)
}

func initLogLevel() {
	switch config.GetString("log.logLevel") {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	}
}
