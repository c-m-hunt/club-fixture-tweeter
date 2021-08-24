# Club Fixture Tweeter

## What you need

- AWS account and key set up
- Play Cricket API key
- `serverless` CLI installed

## Build your config

Place in `clubTweeter/pkg/config/config.json`

```
{
  "playCricket": {
    "clubID": "1234",
    "apiToken": "abc123",
    "teams": [
      "123",
      "456",
      "789"
    ]
  }
  "twitterAuth": {
    "consumerKey": "xxx",
    "consumerSecret": "xxx",
    "accessToken": "xxx",
    "accessSecret": "xxx"
  },
  "twitterMap": {
    "1234": "ClubTwitterWithout@",
  },
  "templates": {
    "fixtures": "Today's Fixtures:\n\n{{.Fixtures}}\n@@mysponsor",
    "fixtureLine": "{{.Team.TeamName}} {{.Venue}} v {{if .Opposition.ClubTwitter}}{{.Opposition.ClubTwitter}}{{else}}{{.Opposition.ClubName}}{{end}}"
  }
}
```

## Run locally

```
go run cli/main.go
```

## Deploy to AWS Lambda

- Amend the cron in `serverless.yml`
- Deploy with: `make deploy`
