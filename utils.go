package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func initPing(server string, token string) {
	ticker := time.NewTicker(time.Second * 10)
	var netClient = &http.Client{
		Timeout: time.Second * 2,
	}
	defer ticker.Stop()

	for range ticker.C {
		req, err := http.NewRequest("GET", "http://"+server+":8080/register", nil)

		if err != nil {
			log.Println(err)
			continue
		}

		req.Header.Set("token", token)
		resp, err := netClient.Do(req)

		if err != nil {
			log.Println(err)
			continue
		}

		bodyByte, e := ioutil.ReadAll(resp.Body)

		if e != nil {
			log.Println(e)
			resp.Body.Close()
			return
		}

		body := string(bodyByte)

		if body == "DONE" {
			log.Println("REGISTER!")
		} else {
			log.Printf("wrong answer from server, body: %s", body)
		}
		resp.Body.Close()
	}
}
