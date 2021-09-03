package gamedata

import (
	"encoding/json"
	"log"

	"github.com/bundgaard/battlenet/hs/gamedata/card"
	"github.com/bundgaard/battlenet/hs/gamedata/cardback"
	"github.com/bundgaard/battlenet/http"
)

type Client struct {
	client *http.Client
}

func NewWith(c *http.Client) *Client {
	return &Client{client: c}
}

func (c *Client) Login() error {
	return c.client.Login()
}

// CARDS

/* CardSearch
endpoint: hearthstone/cards
if search is nil then do regular card search request with rules as laid out in the documentation
else use the search request filtering as entered
*/
func (c *Client) CardSearch(search *card.SearchRequest) {
	url := "/hearthstone/cards/"

	if !search.HasLocale() {
		search.SetLocale("en_US")
	}

	if search != nil {
		url += "?" + search.Encode()
	}

	resp, err := c.client.TokenRequest("GET", nil, url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	http.Todo(resp.Body, "cards%s", search.Encode())

}

// hearthstone/cardbacks
func (c *Client) CardBackSearch(search *cardback.SearchRequest) (*cardback.Response, error) {
	url := "/hearthstone/cardbacks"

	if search != nil {
		if !search.HasLocale() {
			search.SetLocale("en_US")
		}
		url += "?" + search.Encode()
	}

	resp, err := c.client.TokenRequest("GET", nil, url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var ro cardback.Response
	if err := json.NewDecoder(resp.Body).Decode(&ro); err != nil {
		return nil, err
	}

	return &ro, nil
}

// hearthstone/cardbacks/:idorslug
func (c *Client) CardBackBy(idOrSlug, andLocale string) (*cardback.CardBack, error) {
	if andLocale == "" {
		andLocale = "en_US"
	}
	resp, err := c.client.TokenRequest("GET", nil, "/hearthstone/cardbacks/%s?locale=%s", idOrSlug, andLocale)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ro cardback.CardBack
	if err := json.NewDecoder(resp.Body).Decode(&ro); err != nil {
		return nil, err
	}
	return &ro, nil

}

/* // hearthstone/deck
func (c *Client) DeckByCode(code string) {}

// hearthstone/deck
func (c *Client) DeckByList(list []string) {}
*/
// hearthstone/Metadata
func (c *Client) Metadata() {
	resp, err := c.client.TokenRequest("GET", nil, "/hearthstone/metadata/?locale=en_US")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

}

/* // func (c *Client) MetadataSets

// hearthstone/metadata/:type
func (c *Client) MetadataByType(typ string) {}
*/
