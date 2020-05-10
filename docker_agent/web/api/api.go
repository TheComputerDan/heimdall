package api

import (
	"encoding/json"
	"github.com/TheComputerDan/heimdall/docker_agent/docker/connect"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// GetContainers Returns the JSON retrieved from docker.sock
func GetContainers(w http.ResponseWriter, r *http.Request) {
	hostContainers := connect.Containers()
	json.NewEncoder(w).Encode(hostContainers)
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	hostImages := connect.Images()
	json.NewEncoder(w).Encode(hostImages)
}

// Start initiates the API
func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/containers", GetContainers)
	router.HandleFunc("/images", GetImages)
	handlers.AllowedOrigins([]string{"*"})
	// http.ListenAndServe(":8080", router)
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	// handle CORS for local testing purposes
}
