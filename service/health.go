package service

import (
	"fmt"
	"net/http"
)

func InitHealthCheck(){
  
	// ping google.com
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(resp)
	}
}
