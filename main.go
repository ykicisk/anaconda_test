package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"io/ioutil"
)

type ApiConf struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var apiConf ApiConf
	{
		apiConfPath := flag.String("conf", "config.json", "API Config File")
		flag.Parse()
		data, err_file := ioutil.ReadFile(*apiConfPath)
		check(err_file)
		err_json := json.Unmarshal(data, &apiConf)
		check(err_json)
	}

	anaconda.SetConsumerKey(apiConf.ConsumerKey)
	anaconda.SetConsumerSecret(apiConf.ConsumerSecret)
	api := anaconda.NewTwitterApi(apiConf.AccessToken, apiConf.AccessTokenSecret)

	twitterStream := api.PublicStreamSample(nil)
	for {
		x := <-twitterStream.C
		switch tweet := x.(type) {
		case anaconda.Tweet:
			fmt.Println(tweet.Text)
			fmt.Println("-----------")
		case anaconda.StatusDeletionNotice:
			// pass
		default:
			fmt.Printf("unknown type(%T) : %v \n", x, x)
		}
	}
}
