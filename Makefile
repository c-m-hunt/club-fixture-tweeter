.PHONY: build clean deploy

build:
	cd clubTweeter && \
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ./../bin/clubTweeter lambda.go && \
	cd ..
clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy:
	sls deploy --verbose --stage=prod

remove:
	sls remove --verbose --stage=prod