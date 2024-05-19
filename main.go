package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Application struct {
	httpServer *gin.Engine
}

// var configFile = flag.String("f", "details.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()
	if app, err := CreateApp(); err != nil {
		fmt.Printf("error %s", err)
	} else {
		app.httpServer.Run("0.0.0.0:8080")
	}
}
