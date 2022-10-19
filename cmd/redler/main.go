package main

import (
	"flag"
	"fmt"
	"os"

	"redler/internal/pkg/config"
	"redler/internal/pkg/db"
	"redler/internal/router"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// @title           Dashboard Backend API
// @description     redler

// @securityDefinitions.basic  BasicAuth
func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	conf := config.Config{}
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)
		for _, val := range conf.HelpDocs() {
			fmt.Println(val)
		}
		fmt.Println("")
		flag.PrintDefaults()
	}
	flag.Parse()

	err := env.Parse(&conf)
	if err != nil {
		logrus.Error("invalid environment variable: ", err.Error())
		fmt.Printf("%+v\n", errors.WithStack(err))
		return
	}

	logLevel, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		logrus.Error("unable to parse log level: ", err.Error())
		fmt.Printf("%+v\n", errors.WithStack(err))
		return
	}
	logrus.SetLevel(logLevel)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db, err := db.New(&conf)
	if err != nil {
		logrus.Errorf("database verification fail: %v", err.Error())
		fmt.Printf("%+v\n", errors.WithStack(err))
		return
	}

	router.SetupRoutes(r, &conf, db)

	r.Run(fmt.Sprintf(":%s", conf.Port))
}
