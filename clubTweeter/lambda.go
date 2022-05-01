package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/c-m-hunt/club-tweeter/pkg/app"
	"github.com/c-m-hunt/club-tweeter/pkg/img"
)

type Response events.APIGatewayProxyResponse


func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Run successful",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	whichApp, exists := os.LookupEnv("APP")
	if exists {
		if whichApp == "FIX_TWEETER" {
			app.RunFixtureTweet()
		} else if whichApp == "SCORE_UPDATER" {
			app.RunFixtureTweet()
		}
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	// lambda.Start(Handler)
	background := "imgs/backgrounds/landscape.png"
	foreground := "imgs/players/bat/43383.png"
	sponsor := "imgs/ads/PaulRobinson_Logo_1.png"
	output := "result.jpg"

	opts := img.NewCreateLayeredImgOptions()
	sponsorLayer := img.NewImgLayer(sponsor)
	sponsorLayer.X = 30
	sponsorLayer.Y = 30
	sponsorLayer.Scale = 1.6
	layers := img.ImgLayers{
		img.NewImgLayer(foreground),
		sponsorLayer,
	}

	err := img.CreateLayeredImg(background, layers, output, opts)
	if err != nil {
		panic(err)
	}
}
