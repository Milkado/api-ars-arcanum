package helpers

import "github.com/spf13/viper"

func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return viper.GetString(key)
}