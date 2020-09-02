package api

import (
	"encoding/json"
	_ "fmt"
	"github.com/TheComputerDan/sentinel_server/internal/config"
	"github.com/TheComputerDan/sentinel_server/internal/docker/connect"
	"github.com/TheComputerDan/sentinel_server/internal/host"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// GetHost returns info in the form of `Configuration` struct
// defined in the `host.go` file in this project.
func GetHost(w http.ResponseWriter, r *http.Request) {
	var hosts []host.Info

	hostInfo := host.Info{}
	hostInfo.Init()

	hosts = append(hosts, hostInfo)
	json.NewEncoder(w).Encode(hosts)
}

// GetContainers returns the a list of containers.
func GetContainers(w http.ResponseWriter, r *http.Request) {
	hostContainers := connect.Containers()
	json.NewEncoder(w).Encode(hostContainers)
}

// GetImages returns a list of images on the host machine.
func GetImages(w http.ResponseWriter, r *http.Request) {
	hostImages := connect.Images()
	json.NewEncoder(w).Encode(hostImages)
}

//getRESTPort loads the config, and searches for rest_port, defaults to 8096 otherwise
func getRESTPort() string {
	var portNum string

	sConfig := config.Load()
	err := sConfig.ReadInConfig()
	if err != nil {
		//TODO Add more robust error handling
		log.Println("config not found defaulting to port 8096")
		portNum = "8096"
	} else{
		portNum = sConfig.GetString("rest_port")
	}
	return portNum
}

// Start instantiates the API and sets up the endpoints for
// consumption by the server.
func Start() {
	portNum := ":" + getRESTPort()

	router := mux.NewRouter()
	router.HandleFunc("/containers", GetContainers)
	router.HandleFunc("/images", GetImages)
	router.HandleFunc("/host", GetHost)
	handlers.AllowedOrigins([]string{"*"})

	log.Printf("Running server on port %s", portNum)

	http.ListenAndServe(portNum, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	// handle CORS for local testing purposes
}
