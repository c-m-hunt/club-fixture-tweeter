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

type ClubMatches []pc.ClubMatch

func PostTweet() {
	cfg := config.NewConfig()

	c := pc.NewClient(cfg.ClubID, cfg.APIToken)
	fixs := c.GetFixtures(pc.GetCurrentSeason())

	fixToday := ClubMatches(fixs).FilterByDate(time.Now())

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
	} else {
		fmt.Println("No fixtures to send - finding fixtures in next week")
		for i := 1; i < 8; i++ {
			date := time.Now().Add(time.Duration(int(time.Hour) * 24 * i))
			fixFiltered := ClubMatches(fixs).FilterByDate(date)
			if len(fixFiltered) > 0 {
				fmt.Printf("%v fixtures on %v", len(fixFiltered), date.Format("02/01/2006"))
				i = 8
			}
		}
	}
}

func (cms ClubMatches) FilterByDate(d time.Time) ClubMatches {
	fixFiltered := []pc.ClubMatch{}
	date := d.Format("02/01/2006")

	for _, f := range cms {
		if f.Match.MatchDate == date {
			fixFiltered = append(fixFiltered, f)
		}
	}
	fmt.Printf("Filtered from %v fixtures to %v for date %v\n", len(cms), len(fixFiltered), date)
	return fixFiltered
}
