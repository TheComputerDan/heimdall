package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"runtime"
)

const configName = "heimdall_config"

type info struct {
	path      string
	name      string
	extension string
	full      string
}

func initialize() info {

	var configInfo info

	osConfigPath, _ := os.UserConfigDir()

	configInfo.extension = "yaml"
	configInfo.path = fmt.Sprintf("%s/heimdall/", osConfigPath)
	configInfo.name = "heimdall_config"
	configInfo.full = fmt.Sprintf("%s%s.%s", configInfo.path, configInfo.name, configInfo.extension)

	return configInfo
}

func interfaceName() string {
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

// Generate is called to generate a default config file.
func Generate() {
	confInfo := initialize()

	ifaceName := interfaceName()

	agentConf := viper.New()

	agentConf.SetDefault("interface_name", ifaceName)
	agentConf.SetDefault("server_api_port", "8095")
	agentConf.SetDefault("default_polling_interval", "15s")
	agentConf.SetDefault("docker_agent", map[string]string{"enabled": "true", "docker_socket": "/var/run/docker.sock"})

	// Check if the path exists and if not create it
	_, err := os.Stat(confInfo.path)
	if os.IsNotExist(err) {
		os.MkdirAll(confInfo.path, 0755)
	}

	fmt.Println(confInfo.path)

	agentConf.AddConfigPath(confInfo.path)
	agentConf.SetConfigName(confInfo.name)
	agentConf.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	agentConf.SetConfigType(confInfo.extension)
	agentConf.SafeWriteConfig()
}
