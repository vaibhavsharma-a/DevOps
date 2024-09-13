package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/main-page", handleMainPage)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("new-endpoint", handleNewEndpoint)

	addr := "localhost:8080"
	log.Printf("The server is running on the port %s...", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMainPage(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	returnResponse(writer, "This is the main page")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	returnResponse(writer, "All Good!")
}

func handleNewEndpoint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	returnResponse(writer, "This is the new endpoint")
}
func returnResponse(w http.ResponseWriter, responseString string) {
	response := []byte(responseString)
	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
