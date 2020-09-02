package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"

	// "log"
	"os"
	"runtime"
)

const (
	ConfigName = "sentinel_config"
	ConfigType = "yaml"
)

// Configuration is a struct containing config info
type Configuration struct {
	path      string
	name      string
	extension string
	fullname  string
	full      string
}

// Init populates Configuration struct via a pointer
func (conf *Configuration) Init() {
	var (
		configDir  string
	)

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	conf.extension = ConfigType
	conf.path = fmt.Sprintf("%s/sentinel/", configDir)
	conf.name = ConfigName
	conf.fullname = fmt.Sprintf("%s.%s", conf.name, conf.extension)
	conf.full = fmt.Sprintf("%s%s.%s", conf.path, conf.name, conf.extension)

}

// Load returns a viper config loaded and ready to get values from
func Load() *viper.Viper {
	var (
		conf Configuration
		loaded = viper.New()
	)

	conf.Init()

	loaded.AddConfigPath(conf.path)
	loaded.SetConfigName(conf.fullname)
	loaded.SetConfigType(conf.extension)
	err := loaded.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return loaded
}

func defaultInterface() string {
	switch runtime.GOOS {

	case "windows":
		return "Ethernet"

	case "darwin":
		return "en0"

	case "linux":
		return "eth0"

	default:
		return "eth0"

	}
}

// Generate is called to generate a default config file.
func Generate() {
	var (
		conf Configuration
		interfaceName string        = defaultInterface()
		agentConf     *viper.Viper  = viper.New()
	)

	conf.Init()

	agentConf.SetDefault("interface_name", interfaceName)
	agentConf.SetDefault("rest_port", "8096")
	agentConf.SetDefault("grpc_port", "8095")
	agentConf.SetDefault("default_polling_interval", "15s")
	agentConf.SetDefault("docker_agent", map[string]string{"enabled": "true", "docker_socket": "/var/run/docker.sock"})

	// Check if the path exists and if not create it
	_, err := os.Stat(conf.path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(conf.path, 0755)
		if err != nil {
			panic(err)
		}
	}

	agentConf.AddConfigPath(conf.path)
	agentConf.SetConfigName(conf.name)
	_ = agentConf.WriteConfig() // Ignoring Errors to allow to move to SafeWriteConfig
	agentConf.SetConfigType(conf.extension)
	err = agentConf.SafeWriteConfig()
	if err == viper.ConfigFileAlreadyExistsError(conf.full) {
		log.Printf("File already exists continuing...")
	} else if err != nil {
		log.Panic(err) // TODO Consider if this way of error handling will suffice
	}
}
