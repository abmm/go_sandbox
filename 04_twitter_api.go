package main

import (
	"bufio"
	"fmt"
	"github.com/abmm/go_sandbox/Godeps/_workspace/src/github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"time"
)

func start_twitter_api() {
	ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
	ACCESS_TOKEN := os.Getenv("ACCESS_TOKEN")
	CONSUMER_KEY := os.Getenv("CONSUMER_KEY")
	CONSUMER_SECRET := os.Getenv("CONSUMER_SECRET")

	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	fmt.Printf("\nIntroduce the tag to find in Tweeter (max 10 characters):")
	reader := bufio.NewReader(os.Stdin)
	tag, _ := reader.ReadString(byte('\n'))
	if len(tag) > 1 {
		if len(tag) > 10 {
			tag = tag[0:10]
		} else {
			tag = tag[0 : len(tag)-1]
		}
	}
	for {

		//api.PostTweet("probando api", nil)
		v := url.Values{}
		v.Set("count", "10")
		v.Set("lang", "es")
		searchResult, _ := api.GetSearch(string(tag), v)
		i := 0
		for _, tweet := range searchResult.Statuses {
			i++
			//Key =>  "tag/user/:user_id/tweet/:tweet_id/"
			key := tag + "/user/" + tweet.User.IdStr + "/" + "tweet/" + tweet.IdStr
			value := tweet.Text

			addSingleKey(key, value)
		}
		time.Sleep(time.Second * 5)
	}

}
