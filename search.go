package main

import(
	"fmt"
	"net/http"
	"log"
	"net/url"
	"encoding/json"
)

var keywords string = "le OR de OR un OR être OR et OR à OR il OR avoir OR ne OR je OR son OR que OR se OR qui OR ce OR le OR pour OR pas OR que OR vous"
//var keywords string = "test"
//var keywords string = "le OR de OR un OR être OR et OR à OR il OR avoir"
type result struct {
		Statuses []struct {
			Entities struct {
				Hashtags []struct {
					Text string `json:"text"`
				} `json:"hashtags"`
			}
		}
	}

func getTweets() []struct{Text string `json:"text"`}{
	bearer := "Bearer "+getBearer()
	fmt.Println(bearer)

	req,err := http.NewRequest("GET","https://api.twitter.com/1.1/search/tweets.json?q="+url.QueryEscape(keywords),nil)

	req.Header.Add("Host","api.twitter.com")
	req.Header.Add("User-Agent","croisillonzero")
	req.Header.Add("Authorization",bearer)

	resp,err := c.Do(req) // Execute request

	if err != nil {
		log.Fatal(err) // Catch errors
	}
	fmt.Println(resp.Status) // Show success (or not)

	var r result
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatal(err) // Catch errors
	}
	resp.Body.Close()

	fmt.Print("nb result : ")
	nbStatues := len(r.Statuses)
	fmt.Println(nbStatues)
	i := 0
	found := false
	for !found && i<nbStatues {
		found = len(r.Statuses[i].Entities.Hashtags) != 0
		fmt.Println(len(r.Statuses[i].Entities.Hashtags))
		i++
	}
	if !found {log.Fatal("No hastag found")}
	return r.Statuses[i-1].Entities.Hashtags
}


func main() {
	hashtags := getTweets()
	
	fmt.Println("hashtag :"+hashtags[0].Text)
}