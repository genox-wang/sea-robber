package model

import (
	"github.com/sirupsen/logrus"

	"sea-robber-api/app/config"

	"github.com/jinzhu/gorm"
)

var (
	// DB global gorm instance
	DB *gorm.DB
)

// OpenDB open grom
func OpenDB(db string) {
	var err error
	logrus.Info("OpeningDB: ", db)
	DB, err = gorm.Open("mysql", db)

	DB.LogMode(config.GetBool("db.showlog"))

	if err != nil {
		logrus.Fatalf(err.Error())
	}
	Migrate()
}

// Migrate db migration
func Migrate() {
	logrus.Info("Migrate ... ")
	DB.AutoMigrate(
		new(User),
	)
}

// CloseDB close gorm
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
