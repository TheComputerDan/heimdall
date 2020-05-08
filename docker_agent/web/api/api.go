package api

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"heimdall/docker_agent/docker/containers"
	"net/http"
)

// GetContainers Returns the JSON retrieved from docker.sock
func GetContainers(w http.ResponseWriter, r *http.Request) {
	hostContainers := containers.Instantiate()
	// containers.Info(hostContainers)

	json.NewEncoder(w).Encode(hostContainers)
}

// Start initiates the API
func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/containers", GetContainers)
	handlers.AllowedOrigins([]string{"*"})
	// http.ListenAndServe(":8080", router)
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	// handle CORS for local testing purposes
}
