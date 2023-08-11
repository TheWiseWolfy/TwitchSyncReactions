package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type VideoRequest struct {
	streamId string
}

type VideoResponse struct {
	videoUrl string
}

var myMap = map[string]string{
	"U102975477": "https://www.youtube.com/embed/uGGQGoht6ic",
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	var req VideoRequest
	fmt.Printf("REQ %s", string(r.RequestURI))
	req.streamId = r.URL.Query().Get("streamId")
	if req.streamId == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var res VideoResponse
	res.videoUrl = myMap[req.streamId]
	json.NewEncoder(w).Encode(res)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getVideo)
	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":3333", handler)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
