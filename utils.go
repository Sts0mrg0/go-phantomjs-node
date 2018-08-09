package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func initPing(server string, token string) {
	ticker := time.NewTicker(tick)
	netClient := &http.Client{
		Timeout: timeout,
	}
	defer ticker.Stop()

	for range ticker.C {
		err := everyTick(server, token, netClient)

		if err != nil {
			log.Println(err)
			return
		}
	}
}

func everyTick(server string, token string, netClient *http.Client) error {
	req, err := http.NewRequest("GET", fmt.Sprintf(urlHub, server), nil)

	if err != nil {
		log.Println(err)
		return nil
	}

	req.Header.Set(tokenKey, token)
	resp, err := netClient.Do(req)

	if err != nil {
		log.Println(err)
		return nil
	}

	bodyByte, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		resp.Body.Close()
		return e
	}

	body := string(bodyByte)

	if body == done {
		log.Println("REGISTER!")
	} else {
		log.Printf("wrong answer from server, body: %s", body)
	}
	resp.Body.Close()

	return nil
}
