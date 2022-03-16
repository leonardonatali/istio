package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	log.Println("The server is running at port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	log.Fatal(err)
}

func hello(rw http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The message is: %s\n", os.Getenv("MESSAGE"))

	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte(message))
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
