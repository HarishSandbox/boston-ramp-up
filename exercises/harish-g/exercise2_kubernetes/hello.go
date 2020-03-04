package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// main function that prints unix time when queried using
func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		currentTime := time.Now().Unix()
		io.WriteString(w, "Current Unix time is: "+strconv.Itoa(int(currentTime)))
	}

	http.HandleFunc("/hello", helloHandler)
	log.Print(http.ListenAndServe(":8080",nil))
