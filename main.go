package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("start server ...")

	token := os.Getenv("token")

	if token == "" {
		log.Fatalln("no env TOKEN")
		return
	}
	server := os.Getenv("server")

	if server == "" {
		log.Fatalln("no env SERVER")
		return
	}

	go initPing(server, token)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("new request: %s", r.RemoteAddr)
		index(token, w, r)
	})
	err := http.ListenAndServe(":6677", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
