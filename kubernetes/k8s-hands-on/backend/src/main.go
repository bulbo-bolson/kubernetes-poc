package main

import (
	"encoding/json"
	// "fmt"
	// "os"
    "log"
	"net/http"
	"time"
)

type HandsOn struct {
	Time time.Time `json:"time"`
	Hostname string `json:"hostname"`
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn {
		Time: time.Now(),
		Hostname: "localhost",
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResp)
}

func main() {
    s := &server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":9090", nil))
}