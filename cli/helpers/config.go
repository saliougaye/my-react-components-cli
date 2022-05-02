package helpers

import "github.com/spf13/viper"

func SaveInConfigFile(key string, value interface{}) error {

	viper.Set(key, value)

	err := viper.WriteConfig()

	return err
}
