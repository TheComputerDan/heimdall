package api

import (
	"encoding/json"
	"fmt"
	"github.com/TheComputerDan/heimdall/docker_agent/docker/connect"
	"github.com/TheComputerDan/heimdall/docker_agent/host"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

func GetHost(w http.ResponseWriter, r *http.Request) {
	hosts := []host.Info{}
	rb := host.HostInfo()
	hosts = append(hosts, rb)
	json.NewEncoder(w).Encode(hosts)
}

// GetContainers Returns the JSON retrieved from docker.sock
func GetContainers(w http.ResponseWriter, r *http.Request) {
	hostContainers := connect.Containers()
	json.NewEncoder(w).Encode(hostContainers)
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	hostImages := connect.Images()
	json.NewEncoder(w).Encode(hostImages)
}

func loadConfig() string {
	viper.SetConfigName("heimdall.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("docker_agent/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	portNum := viper.GetString("heimdall_port")
	return portNum
}

// Start initiates the API
func Start() {
	portNum := ":" + loadConfig()

	router := mux.NewRouter()
	router.HandleFunc("/containers", GetContainers)
	router.HandleFunc("/images", GetImages)
	router.HandleFunc("/host", GetHost)
	handlers.AllowedOrigins([]string{"*"})
	// http.ListenAndServe(":8080", router)
	http.ListenAndServe(portNum, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	// handle CORS for local testing purposes
}
