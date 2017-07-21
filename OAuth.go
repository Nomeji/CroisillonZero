package main

import(
	"fmt"
	"net/http"
	"net/url"
	"io"
	"strings"
	"log"
	"encoding/base64"
	"encoding/json"
)

var c http.Client
var r io.Reader = strings.NewReader("grant_type=client_credentials")

func main() {
	//Formatting autorization
	consumerKey := url.QueryEscape("6BfRu5tBKz6cRXZK6vro1g")//URL encode
	consumerSecret := url.QueryEscape("YOG53ipiTisvq4KZ6VFZKrDR5ommLrGGoDhsE06GA")

	concat := []byte(consumerKey+":"+consumerSecret)//Concat keys
	fmt.Println("concat")

	auth := "Basic "+base64.StdEncoding.EncodeToString(concat)//Base64 encode keys
	fmt.Println(auth)	


	req,err := http.NewRequest("POST","https://api.twitter.com/oauth2/token",r) // Creating the request

	req.Header.Add("Authorization",auth)
	req.Header.Add("Host","api.twitter.com")
	req.Header.Add("User-Agent","croisillonzero")
	req.Header.Add("Content-Type","application/x-www-form-urlencoded;charset=UTF-8")

	resp,err := c.Do(req) // Execute request

	if err != nil {
		log.Fatal(err) // Catch errors
	}
	fmt.Printf(resp.Status) // Show success (or not)

	// Getting the JSON
	var data struct {Token string}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err) // Catch errors
	}
	bearer := data.Token
	fmt.Printf(bearer)

	/*type s struct = { Key, Value string}
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&s)
	bearer := */

}