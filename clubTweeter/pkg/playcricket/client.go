package playcricket

type Client struct {
	SiteID   string
	APIToken string
}

func NewClient(siteID string, apiToken string) Client {
	return Client{SiteID: siteID, APIToken: apiToken}
}
