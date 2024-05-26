package config

import (
	"fmt"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// path string injected by app/main.go
func New(path string) (*viper.Viper, error) {
	fmt.Printf("====== path ====== %s", path)

	var (
		err error
		v   = viper.New()
	)
	v.AddConfigPath(".")
	v.SetConfigFile(string(path))

	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	return v, err
}

var ProviderSet = wire.NewSet(New)
