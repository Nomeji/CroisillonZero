package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"math/rand"
	"time"
)

var c http.Client
var requestParams io.Reader = strings.NewReader("grant_type=client_credentials")
var consumerKey string = "6BfRu5tBKz6cRXZK6vro1g"
var consumerSecret string = "YOG53ipiTisvq4KZ6VFZKrDR5ommLrGGoDhsE06GA"

func getBearer() string {
	//Formatting autorization
	key := url.QueryEscape(consumerKey) //URL encode
	secret := url.QueryEscape(consumerSecret)

	concat := []byte(key + ":" + secret) //Concat keys

	auth := "Basic " + base64.StdEncoding.EncodeToString(concat) //Base64 encode keys
	fmt.Println(auth)

	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", requestParams) // Creating the request

	req.Header.Add("Authorization", auth)
	req.Header.Add("Host", "api.twitter.com")
	req.Header.Add("User-Agent", "croisillonzero")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := c.Do(req) // Execute request

	if err != nil {
		log.Fatal(err) // Catch errors
	}
	fmt.Println(resp.Status) // Show success (or not)

	// Getting the JSON
	var data struct {
		Token string `json:"access_token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err) // Catch errors
	}
	resp.Body.Close()
	bearer := data.Token
	return bearer
}

func genNonce() string {
	t := time.Now()
	r := rand.New(rand.NewSource(t.Unix()))
	b := make([]byte, 32)
	_,err := io.ReadFull(r,b)
	if err != nil {
		log.Fatal(err) // Catch errors
	}
	b64rand := base64.StdEncoding.EncodeToString([]byte(b))
	return b64rand
}

func threeLegged() string {

	autorization := "OAuth "+
					"oauth_consumer_key=\""+url.QueryEscape(consumerKey)+"\"," +
					"oauth_nonce=\""+url.QueryEscape(genNonce())+"\""+
					"oauth_signature=\"\"," +
					"oauth_signature_method=\"\"," +
					"oauth_timestamp=\"\"," +
					"oauth_token=\"\"," +
					"oauth_version=\"1.0\""

	return autorization
}