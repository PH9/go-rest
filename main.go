package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Print("[I] API is starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	fmt.Println(" on port " + port)

	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{"message":"hello world"}`)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}

	fmt.Fprintf(w, "%q", body)
}
