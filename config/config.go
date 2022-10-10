package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig(c_path string, c_name string, c_type string) (config *viper.Viper) {
	config = viper.New()
	config.AddConfigPath(c_path)
	config.SetConfigName(c_name)
	config.SetConfigType(c_type)
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	return config
}
func AutoConfig() (config *viper.Viper) {
	c_path := flag.String("path", "", "file path")
	c_name := flag.String("name", "", "file name")
	c_type := flag.String("type", "", "file type")
	flag.Parse()

	return GetConfig(*c_path, *c_name, *c_type)
}
