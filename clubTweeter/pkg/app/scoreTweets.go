package app

import (
	_ "embed"
	"image/color"
	"math"

	"github.com/c-m-hunt/club-tweeter/pkg/config"
	"github.com/c-m-hunt/club-tweeter/pkg/img"
	pc "github.com/c-m-hunt/club-tweeter/pkg/playcricket"
	"github.com/fogleman/gg"
)

type ScoreCache struct {
	MatchId string `json:"matchId"`
	Fifties []string `json:"fifties"`
	Hundreds []string `json:"hundreds"`
	FiveWickets []string `json:"fiveWickets"`
}

var cfg config.Config

func init() {
	cfg = config.NewConfig()
}

func RunScoreTweet() error {
	c := pc.NewClient(cfg.PlayCricket.ClubID, cfg.PlayCricket.APIToken)
	_ = pc.ClubMatches(c.GetFixtures(pc.GetCurrentSeason()))


	return nil
	// Get live fixtures

	// Find any new 50s or 5 fors

	// Save 50s or 5 fors

	// Tweet score updates
}

type ScoreEventType int64

const (
	BatFifty ScoreEventType = iota
	BatHundred
	BowlFiveWickets
)

type ScoreEvent struct {
	PlayerId int
	ScoreEventType
}

func (se *ScoreEvent) GenerateImage(filePath string) error {
	background := "imgs/backgrounds/landscape.png"
	foreground := "imgs/players/ball/1417520.png"
	sponsor := "imgs/ads/PaulRobinson_Logo_1.png"

	sponsorLayer := img.NewImgLayer(sponsor)
	sponsorLayer.X = 30
	sponsorLayer.Y = 30
	sponsorLayer.Scale = 1.6

	foregroundLayer := img.NewImgLayer(foreground)
	foregroundLayer.X = -150

	var scoreText string
	if se.ScoreEventType == BatFifty {
		scoreText = cfg.ScoreImgs.FiftyText
	} else if se.ScoreEventType == BatHundred {
		scoreText = cfg.ScoreImgs.HundredText
	} else if se.ScoreEventType == BowlFiveWickets {
		scoreText = cfg.ScoreImgs.FiveWicketsText
	}

	imgBack, err := gg.LoadImage(background)
	if err != nil {
		return err
	}

	layers := img.Layers{
		foregroundLayer,
		sponsorLayer,
		img.TextLayer{"A Wainwright", 50, imgBack.Bounds().Dy() - 50, color.White, 0, 150, false},
		img.TextLayer{
			scoreText,
			imgBack.Bounds().Dy() / 2,
			-1 * imgBack.Bounds().Dx() + 150,
			color.White, math.Pi / 2,
			400,
			true,
		},
	}

	return img.CreateLayeredImg(background, layers, filePath)
}