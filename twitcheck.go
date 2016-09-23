// twitcheck is a simple command-line tool to pull data on a single twitter user.
// Just for it it also finds their most commonly used word in the last five tweets.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

//LoadCredentials reads a json config file that contains the ConsumerKey & ConsumerSecret.
//It uses them to build an application authorized Twitter API client.
func LoadCredentials() (client *twittergo.Client, err error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)
	config := &oauth1a.ClientConfig{}
	err = dec.Decode(config)
	if err != nil {
		return nil, err
	}
	client = twittergo.NewClient(config, nil)
	return client, nil
}

//printResponseInfo prints status information related to the last call out the Twitter API
func printResponseInfo(resp *twittergo.APIResponse, timeElapsed time.Duration) {
	fmt.Println("---------------------Info--------------------")
	fmt.Println("Status:              ", resp.Status)
	fmt.Printf("Time Elapsed:         %fs\n", timeElapsed.Seconds())
	if resp.HasRateLimit() {
		fmt.Println("Rate limit:          ", resp.RateLimit())
		fmt.Println("Rate limit remaining:", resp.RateLimitRemaining())
		fmt.Println("Rate limit reset:    ", resp.RateLimitReset())
	} else {
		fmt.Println("(Could not parse rate limit from response.)")
	}
}

//Histogram is a Absolute Frequency table for words.
type Histogram struct {
	table map[string]int
}

//NewHistogram returns a reference to a newly instanced Histogram.
func NewHistogram() *Histogram {
	return &Histogram{table: make(map[string]int)}
}

//AddWords normalizes text to lower case and then breaks it down into words that add added to the Histogram
func (h *Histogram) AddWords(text string) {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	for _, word := range words {
		v := h.table[word]
		h.table[word] = v + 1
	}
}

//MostCommonWord finds the most frequently occurring word in the Histogram.
func (h *Histogram) MostCommonWord() string {
	mcWord := ""
	max := 0
	for word, count := range h.table {
		if count > max {
			mcWord = word
			max = count
		}
	}
	return mcWord
}

var (
	whitespaceRegExp = regexp.MustCompile("\\s+")
)

//normalizeWhiteSpace collapses all contiguous whitespace into a single space
func normalizeWhiteSpace(text string) string {
	return whitespaceRegExp.ReplaceAllString(text, " ")
}

func main() {
	//Instance a Twitter API client using Credentials file.
	client, err := LoadCredentials()
	if err != nil {
		log.Fatal("Could not parse config.json file:", err.Error())
	}
	//Collect the ScreenName Argument.
	var screenName string
	if 2 > len(os.Args) {
		log.Fatal("No Twitter ScreenName provided. \nUsage: twitcheck [screen_name]\n")
	} else {
		screenName = os.Args[1]
	}
	//Build & send the Twitter API Query to search for the Supplied ScreenName.
	query := url.Values{}
	query.Set("screen_name", screenName)
	uri := fmt.Sprintf("/1.1/users/show.json?%v", query.Encode())
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal("Could not parse request: ", err.Error())
	}
	t := time.Now()
	resp, err := client.SendRequest(req)
	timeElapsed := time.Since(t)
	if err != nil {
		log.Fatal("Could not send request: ", err.Error())
	}
	//Parse the JSON response into a User struct that provides some nice helper methods on top of a map[string]interface{}.
	var user twittergo.User
	err = resp.Parse(&user)
	if err != nil {
		log.Fatal("Problem parsing response: ", err.Error())
	}
	//Print Data from User's Profile and Info about the response.
	fmt.Println("-------------------Profile-------------------")
	fmt.Println("Name:                ", user.Name())
	fmt.Println("Screen Name:         ", user.ScreenName())
	fmt.Println("Description:         ", user["description"].(string))
	fmt.Println("Location:            ", user["location"].(string))
	fmt.Println("Tweets:              ", user["statuses_count"].(float64))
	fmt.Println("Followers:           ", user["followers_count"].(float64))
	fmt.Println("Following:           ", user["friends_count"].(float64))
	fmt.Println()
	printResponseInfo(resp, timeElapsed)
	//Build & Send a Twitter API query for the 5 most recent tweets from the User's timeline.
	query = url.Values{}
	query.Set("id", user.IdStr())
	query.Set("count", "5")
	uri = fmt.Sprintf("/1.1/statuses/user_timeline.json?%v", query.Encode())
	req, err = http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal("Could not parse request: ", err.Error())
	}
	t = time.Now()
	resp, err = client.SendRequest(req)
	timeElapsed = time.Since(t)
	if err != nil {
		log.Fatal("Could not send request: ", err.Error())
	}
	//Parse the JSON response into a slice of Tweets
	var tweets []twittergo.Tweet
	err = resp.Parse(&tweets)
	if err != nil {
		log.Fatal("Problem parsing response: ", err.Error())
	}
	//Print any tweets received if there were any & any response information.
	fmt.Println("\n-------------------Tweets-------------------")
	if len(tweets) > 0 {
		histogram := NewHistogram() //Build a histogram of words to find the Most Common word this users likes to use.
		for i, tweet := range tweets {
			fmt.Println(i+1, ":", normalizeWhiteSpace(tweet.Text()))
			fmt.Println("\tBy,", tweet.User().Name(), "at", tweet.CreatedAt().Local())
			fmt.Println()
			histogram.AddWords(tweet.Text())
		}
		fmt.Println("Most Common Word Tweeted:", histogram.MostCommonWord())
	} else {
		fmt.Println("(Could not collect Tweets.)")
	}
	printResponseInfo(resp, timeElapsed)
}
