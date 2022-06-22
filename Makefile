.PHONY: build clean deploy

build:
	cd clubTweeter && \
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ./../bin/clubTweeter lambda.go && \
	cd ..

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy:
	make build
	export AWS_PROFILE=personal
	sls deploy --verbose --stage=prod

remove:
	sls remove --verbose --stage=prod

sync_files:
	AWS_PROFILE=personal aws s3 sync ./clubTweeter/imgs s3://sosemtcc/imgs