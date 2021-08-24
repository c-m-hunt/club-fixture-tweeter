package app

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/c-m-hunt/club-tweeter/pkg/config"
	pc "github.com/c-m-hunt/club-tweeter/pkg/playcricket"
	"github.com/c-m-hunt/club-tweeter/pkg/twitter"
)

func PostTweet() {
	cfg := config.NewConfig()

	c := pc.NewClient(cfg.PlayCricket.ClubID, cfg.PlayCricket.APIToken)
	fixs := pc.ClubMatches(c.GetFixtures(pc.GetCurrentSeason()))
	fixs.PopulateTwitter(cfg.TwitterMap)

	for i := 0; i < 8; i++ {
		date := time.Now().Add(time.Duration(int(time.Hour) * 24 * i))
		fixFiltered := pc.ClubMatches(fixs).FilterByDate(date, cfg.PlayCricket.Teams)
		if len(fixFiltered) > 0 {
			tweetStr := twitter.GenerateFixtureTweet(fixFiltered, cfg)
			fmt.Printf("----------------------\n%v\n----------------------\n", tweetStr)
			// if i == 0 {
			// 	twitter.SendTweet(cfg.TwitterAuth, tweetStr)
			// }
		}
	}
}
