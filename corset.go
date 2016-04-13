package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if len(url) == 0 {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	res, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		http.Error(w, msg, http.StatusBadGateway)
		return
	}

	if err != nil {
		msg := fmt.Sprintf("%v", err)
		http.Error(w, msg, http.StatusBadGateway)
		return
	}
	for key := range res.Header {
		w.Header().Set(key, res.Header.Get(key))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.Copy(w, res.Body)

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	http.HandleFunc("/", handler)
	fmt.Println("http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
