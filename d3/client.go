package d3

import (
	"github.com/bundgaard/battlenet/d3/community"
	"github.com/bundgaard/battlenet/d3/gamedata"
	"github.com/bundgaard/battlenet/http"
)

type Client struct {
	client    *http.Client
	GameData  *gamedata.GameData
	Community *community.Community
}

func New(clientId, clientSecret, region string) *Client {
	c := http.New(clientId, clientSecret, region)
	return &Client{
		client:    c,
		GameData:  gamedata.NewWith(c),
		Community: community.New(c),
	}
}

func (c *Client) Login() error {
	return c.client.Login()
}

func (c Client) Media(mt http.MediaType, size string, icon string) ([]byte, error) {
	return c.client.Media(mt, size, icon)
}
