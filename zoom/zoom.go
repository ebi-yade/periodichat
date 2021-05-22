package zoom

import (
	"github.com/go-resty/resty/v2"
)

var (
	BaseURL = "https://zoom.us/"
)

type Client struct {
	token    string
	endpoint string
	rClient  *resty.Client
}

func New(token string) *Client {
	return &Client{
		token:    token,
		endpoint: BaseURL,
		rClient:  resty.New(),
	}
}
