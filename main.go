package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var port string

func main() {
	fmt.Print("[I] API is starting...")

	port = os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	fmt.Println(" on port " + port)

	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{"message":"hello world"}`)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, `{"message":"cannot get body"}`)
		return
	}

	res := defaultResponse{Port: port, Path: r.URL.Path, Body: string(body)}
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, `{"message":"cannot parse json"}`)
		return
	}

	fmt.Fprint(w, string(b))
}

type defaultResponse struct {
	Port string
	Path string
	Body string
}
