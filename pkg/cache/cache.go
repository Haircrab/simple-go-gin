package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Options struct {
	Url string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("redis", o); err != nil {
		return nil, err
	}

	return o, err
}

func New(o *Options) *redis.Client {
	redis_uri := o.Url

	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_uri,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// test connection
	if err := rdb.Set("key", "value", 10*time.Second).Err(); err != nil {
		panic(err)
	}
	if val, err := rdb.Get("key").Result(); err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
	}

	return rdb
}

var ProviderSet = wire.NewSet(NewOptions, New)
