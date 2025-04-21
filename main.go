package main

import (
	"GST/billlingSystem/config"
	"GST/billlingSystem/routes"
	"GST/billlingSystem/service"
	"fmt"
	"strconv"

	logger "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	config.Load()
	startApp()
}

func startApp() (err error) {
	deps, err := service.InitDependencies()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Database init failed")
		return
	}

	router := routes.InitRouter(deps)

	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort()
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
	return
}
