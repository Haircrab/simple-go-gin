package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

type healthCheck struct {
	Status string `json:"status"`
}

func pingHealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, healthCheck{Status: "OK"})
}


func main() {
	// ping redis
	redis_uri := os.Getenv("REDIS_URI")


  
	// log.Printf("Connecting to redis at %s", redis_uri)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_uri,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set( "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get( "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	val2, err := rdb.Get( "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// ping google.com
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(resp)
	}

	// set up gin
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/", pingHealthCheck)

	router.Run("0.0.0.0:8080")
}
