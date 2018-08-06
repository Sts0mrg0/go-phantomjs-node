package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	http.HandleFunc("/phantom", func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse("http://127.0.0.1:4444")
		tokenEnv := os.Getenv("token")
		token := r.URL.Query().Get("token")

		if token != tokenEnv {
			w.WriteHeader(400)
			w.Write([]byte("bad token"))
			log.Println("bad token")
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
