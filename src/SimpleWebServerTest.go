package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serve)
	error := http.ListenAndServe("localhost:8080", nil)
	if error != nil {
		log.Fatal("server start failed!")
	}
}

func serve(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	for k, v := range request.Form {
		fmt.Println(k, v)
	}
	fmt.Fprintf(writer, "Hello web!")
}
