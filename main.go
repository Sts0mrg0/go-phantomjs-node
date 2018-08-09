package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

var (
	badToken   = "bad token"
	tick       = time.Second * 10
	timeout    = time.Second * 2
	pong       = "PONG"
	done       = "DONE"
	tokenKey   = "token"
	serverKey  = "server"
	noEnv      = "no env %s"
	newRequest = "new request %s: %s"
	urlHub     = "http://%s:8080/register"
)

func main() {
	log.Println("start server ...")

	token := os.Getenv(tokenKey)

	if token == "" {
		log.Fatalf(noEnv, tokenKey)
		return
	}
	server := os.Getenv(serverKey)

	if server == "" {
		log.Fatalf(noEnv, serverKey)
		return
	}

	go initPing(server, token)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(newRequest, r.RequestURI, r.RemoteAddr)
		index(token, w, r)
	})

	u, _ := url.Parse("http://127.0.0.1:4444")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(newRequest, r.RequestURI, r.RemoteAddr)
		tokenEnv := os.Getenv(tokenKey)
		token := r.Header.Get(tokenKey)

		if token != tokenEnv {
			w.WriteHeader(400)
			w.Write([]byte(badToken))
			log.Println(badToken)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":6677", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
