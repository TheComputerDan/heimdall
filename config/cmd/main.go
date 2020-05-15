package main

import (
	"github.com/spf13/viper"
)

func writeDefault() {

}

func main() {

	viper.SetDefault("interface_name", "en0")
	viper.SetDefault("heimdall_port", "8095")
	// viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.AddConfigPath("cfg/")
	viper.SetConfigName("config")
	viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	viper.SetConfigType("yaml")
	viper.SafeWriteConfig()
	// viper.WriteConfigAs("cfg/config.yaml")

}
