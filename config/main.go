package main

import (
	"github.com/spf13/viper"
	"runtime"
)

func defaultInterfaceName() string {
	osType := runtime.GOOS

	switch osType {
	case "windows":
		return "Ethernet"
	case "darwin":
		return "en0"
	case "linux":
		return "eth0"
	default:
		return "eth0" // Figure out a better default
	}

}

func Generate() {
	ifaceName := defaultInterfaceName()

	viper.SetDefault("interface_name", ifaceName)
	viper.SetDefault("server_api_port", "8095")

	viper.AddConfigPath("../docker_agent/config")
	viper.SetConfigName("agent")
	viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	viper.SetConfigType("yaml")
	viper.SafeWriteConfig()
	// viper.WriteConfigAs("cfg/config.yaml")
}

func main() {
	Generate()
}
