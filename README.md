# Club Fixture Tweeter

## What you need

- AWS account and key set up
- Play Cricket API key
- `serverless` CLI installed

## Build your config

Place in `clubTweeter/pkg/config/config.json`

```
{
  "clubID": "1234",
  "apiToken": "abc123",
  "teams": [
    "123",
    "456",
    "789"
  ],
  "twitterAuth": {
    "consumerKey": "xxx",
    "consumerSecret": "xxx",
    "accessToken": "xxx",
    "accessSecret": "xxx"
  },
  "twitterMap": {},
  "templates": {
    "fixtures": "Today's Fixtures:\n\n{{.Fixtures}}\n@@mysponsor",
    "fixtureLine": "{{.Team.TeamName}} {{.Venue}} v {{.Opposition.ClubName}}"
  }
}
```

## Run locally

```
go run cli/main.go
```

## Deploy to AWS Lambda

- Amend the cron in `serverless.yml`
- Run `make build`
- Deploy with:

```
make deploy
```
