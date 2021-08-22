```
cd clubTweeter && \
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ./../bin/clubTweeter main.go && \
cd ..
```
