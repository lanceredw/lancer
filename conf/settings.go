package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitSettings() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Load Config Error: %s", err.Error()))
	}
}
