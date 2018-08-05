package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func initPing(server string, token string) {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for range ticker.C {
		var netClient = &http.Client{
			Timeout: time.Second * 2,
		}
		resp, e := netClient.Get("http://" + server + ":8080/register?token=" + token)

		if e != nil {
			log.Println(e)
			continue
		}

		defer resp.Body.Close()

		bodyByte, e := ioutil.ReadAll(resp.Body)

		if e != nil {
			log.Println(e)
			return
		}

		body := string(bodyByte)

		if body == "DONE" {
			log.Println("REGISTER!")
		} else {
			log.Printf("wrong answer from server, body: %s", body)
		}
	}
}
