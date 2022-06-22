package playcricket

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	SiteID   string
	APIToken string
	HttpClient *http.Client
}

func NewClient(siteID string, apiToken string) Client {
	client := Client{SiteID: siteID, APIToken: apiToken}
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	client.HttpClient = &httpClient
	return client
}

func (c Client) getData(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := c.HttpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	return body, nil
}
