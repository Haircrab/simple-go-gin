package main

import (
	"crab-dev/simple-go-gin/service"
)

func main() {
	service.InitDistributedCache()

	service.InitHealthCheck()

	service.InitRouter()
}
