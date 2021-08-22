package app

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/c-m-hunt/club-tweeter/pkg/config"
	pc "github.com/c-m-hunt/club-tweeter/pkg/playcricket"
	"github.com/c-m-hunt/club-tweeter/pkg/twitter"
	twt "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func PostTweet() {
	cfg := config.NewConfig()

	c := pc.NewClient(cfg.ClubID, cfg.APIToken)
	fixs := c.GetFixtures(pc.GetCurrentSeason())

	date := time.Now().Format("02/01/2006")
	fmt.Print(date)
	fixToday := []pc.ClubMatch{}

	for _, f := range fixs {
		if f.Match.MatchDate == date {
			fixToday = append(fixToday, f)
		}
	}

	if len(fixToday) > 0 {
		tweetStr := twitter.GenerateFixtureTweet(fixToday, cfg)
		fmt.Printf("%+v\n", cfg)
		// https://github.com/dghubble/oauth1
		twtCfg := oauth1.NewConfig(cfg.TwitterAuth.ConsumerKey, cfg.TwitterAuth.ConsumerSecret)
		token := oauth1.NewToken(cfg.TwitterAuth.AccessToken, cfg.TwitterAuth.AccessSecret)
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

}
