package main

import (
	"net/http"
	"log"
)

func index(tokenEnv string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token := r.URL.Query().Get("token")

	if token != tokenEnv {
		w.WriteHeader(400)
		w.Write([]byte("bad token"))
		log.Println("bad token")
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("PONG"))
	log.Println("PONG")
}