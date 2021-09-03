package gamedata

import (
	"encoding/json"

	"github.com/bundgaard/battlenet/http"
)

func NewWith(client *http.Client) *GameData {
	return &GameData{client: client}
}

type GameData struct {
	client *http.Client
}

// SeasonIndex /data/d3/season
func (gd GameData) SeasonIndex() (interface{}, error) {
	resp, err := gd.client.TokenRequest("GET", nil, "/data/d3/season/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var v interface{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// Season /data/d3/season/:id
func (gd GameData) Season(id int) (interface{}, error) {
	resp, err := gd.client.TokenRequest("GET", nil, "/data/d3/season/%d", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var v interface{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// Leaderboard /data/d3/season/:id/leaderboard/:leaderboard
func (gd GameData) Leaderboard(id int, leaderboard string) (interface{}, error) {
	resp, err := gd.client.TokenRequest("GET", nil, "/data/d3/season/%d/leaderboard/%s", id, leaderboard)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var v interface{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// /data/d3/era
func (gd GameData) EraIndex() (interface{}, error) {
	resp, err := gd.client.TokenRequest("GET", nil, "/data/d3/era/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var v interface{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// /data/d3/era/:id
func (gd GameData) Era(id int) (interface{}, error) {
	resp, err := gd.client.TokenRequest("GET", nil, "/data/d3/era/%d", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var v interface{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// /data/d3/era/:id/leaderboard/:leaderboard
func (gd GameData) EraLeaderBoard(id int, leaderboard string) (interface{}, error) {
	resp, err := gd.client.TokenRequest("GET", nil, "/data/d3/era/%d/leaderboard/%s", id, leaderboard)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var v interface{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func (gd GameData) Login() error {
	return gd.client.Login()
}
