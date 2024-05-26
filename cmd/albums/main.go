package main

import (
	"flag"
	"fmt"
)

var configFile = flag.String("f", "../../configs/env.yaml", "set config file which viper will loading.")

func main() {
	fmt.Println("hello from services")
	flag.Parse()

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
