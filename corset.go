package main

import (
	"fmt"
	"io"
	"net/http"
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
	fmt.Printf("HEADERS %v", res.Header)
	for key := range res.Header {
		w.Header().Set(key, res.Header.Get(key))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.Copy(w, res.Body)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
