package twitter

import (
	"fmt"

	"github.com/c-m-hunt/club-tweeter/pkg/config"
	twt "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func SendTweet(auth config.TwitterAuth, tweetStr string) {
	// https://github.com/dghubble/oauth1
	twtCfg := oauth1.NewConfig(auth.ConsumerKey, auth.ConsumerSecret)
	token := oauth1.NewToken(auth.AccessToken, auth.AccessSecret)
	httpClient := twtCfg.Client(oauth1.NoContext, token)
	client := twt.NewClient(httpClient)
	// https://github.com/dghubble/go-twitter
	tweet, resp, err := client.Statuses.Update(tweetStr, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(tweet)
}
