package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

func getTweet() result {
	req,err := http.NewRequest("GET","https://stream.twitter.com/1.1/statuses/sample.json",nil)
	req.Header.Add("Authorization",bearer)
	req.Header.Add("Host","stream.twitter.com")
	req.Header.Add("User-Agent","croisillonzero")

	resp,err := c.Do(req) // Execute request

	if err != nil {
		log.Fatal(err) // Catch errors
	}
	fmt.Println("stream status : "+resp.Status) // Show success (or not)

	var r result
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatal(err) // Catch errors
	}
	resp.Body.Close()
	fmt.Print("nb result : ")
	fmt.Println(len(r.Statuses))
	return r
}