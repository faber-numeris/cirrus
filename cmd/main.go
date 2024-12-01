package main

import (
	"github.com/faber-numeris/cirrus/internal/routes"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   false,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var rtr = routes.NewRouter()
	rtr.Run()

}
