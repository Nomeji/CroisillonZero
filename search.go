package main

import(
	"fmt"
	"net/http"
	"log"
	"net/url"
	"encoding/json"
	"errors"
)

var frenchWords string = "le OR de OR un OR être OR et OR à OR il OR avoir OR ne OR je OR son OR que OR se OR qui OR ce OR le OR pour OR pas OR que OR vous"
var bearer string

type result struct {
		Statuses []struct {
			Id string `json:"id_str"`
			Entities struct {
				Hashtags []struct {
					Text string `json:"text"`
				} `json:"hashtags"`
			}
			User struct {
				Screen_name string `json:"screen_name"`
			} `json:"user"` 
		}
	}

func search(keywords string,maxId string) result {
	//fmt.Println(bearer)
	var req *http.Request
	var err error
	if maxId == "" {
		req,err = http.NewRequest("GET","https://api.twitter.com/1.1/search/tweets.json?lang=fr&q="+url.QueryEscape(keywords),nil)
		fmt.Println("Issued request: "+"https://api.twitter.com/1.1/search/tweets.json?lang=fr&q="+url.QueryEscape(keywords))
	}else{
		req,err = http.NewRequest("GET","https://api.twitter.com/1.1/search/tweets.json?lang=fr&q="+url.QueryEscape(keywords)+"&max_id="+url.QueryEscape(maxId),nil)
		fmt.Println("Issued request: "+"https://api.twitter.com/1.1/search/tweets.json?lang=fr&q="+url.QueryEscape(keywords)+"&max_id="+url.QueryEscape(maxId))
	}
	req.Header.Add("Host","api.twitter.com")
	req.Header.Add("User-Agent","croisillonzero")
	req.Header.Add("Authorization",bearer)

	resp,err := c.Do(req) // Execute request

	if err != nil {
		log.Fatal(err) // Catch errors
	}
	fmt.Println("search status : "+resp.Status) // Show success (or not)

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

func getHastags(maxId string) ([]string, string) {
	r := search(frenchWords,maxId)
	
	nbStatues := len(r.Statuses)
	found := false
	var hashtags [] string
	for i:=0;i<nbStatues;i++ { // browse trough statues
		var nbHashtags = len(r.Statuses[i].Entities.Hashtags)
		found = nbHashtags != 0
		if found {
			for j:=0;j<nbHashtags;j++ { // browse trough hashtag(s) of statues
				hashtags = append(hashtags,r.Statuses[i].Entities.Hashtags[j].Text)
			}
		}
	}
	lastId := r.Statuses[nbStatues-1].Id
	return hashtags,lastId
}

func verifHashtags(hashtags []string) (result,error) {
	unique := false
	size := len(hashtags)
	i := 0
	var r result

	for !unique && i < size {
		r = search("#"+hashtags[i],"") //Doesn't need max id so set to ""
		unique = len(r.Statuses) == 1 // Is the hashtag unique?
		i++
	}
	if unique {
		return r,nil
	}else{
		return r,errors.New("not unique") //r is nil if no unique hashtag is found
	}
}


func main() {
	/*bearer = "Bearer "+getBearer()

	tweet := getTweet()
	fmt.Println(tweet.Statuses[0].Id)*/

	/*var uniqueTweet result
	var hashtags []string
	var maxId string = ""
	var err = errors.New("not unique")

	for err != nil {
		hashtags,maxId = getHastags(maxId)
		for i := 0; i < len(hashtags); i++ {
			fmt.Println("hashtag : "+hashtags[i])
		}
		uniqueTweet,err = verifHashtags(hashtags)
	}

	screen_name := uniqueTweet.Statuses[0].User.Screen_name
	id := uniqueTweet.Statuses[0].Id
	fmt.Print("https://twitter.com/"+screen_name+"/status/")
	fmt.Println(id)*/

	test := genNonce()
	fmt.Println(test)
	fmt.Println(threeLegged())
}