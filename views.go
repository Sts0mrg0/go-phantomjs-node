package main

import (
	"log"
	"net/http"
)

func index(tokenEnv string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get(tokenKey)

	if token != tokenEnv {
		w.WriteHeader(400)
		w.Write([]byte(badToken))
		log.Println(badToken)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(pong))
	log.Println(pong)
}
