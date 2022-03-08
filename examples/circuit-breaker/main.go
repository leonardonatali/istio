package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ":80"

	log.Printf("the server is running at port %s", port)

	if err := http.ListenAndServe(port, http.HandlerFunc(handle)); err != nil {
		log.Fatal(err)
	}
}

func handle(rw http.ResponseWriter, r *http.Request) {
	if os.Getenv("break") != "" {
		time.Sleep(10 * time.Second)
		rw.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("ok :)\n"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	return
}
