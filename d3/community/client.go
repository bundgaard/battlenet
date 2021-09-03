// Copyright (C) 2021 David Bundgaard
package community

import (
	"encoding/json"

	"github.com/bundgaard/battlenet/d3/community/types"
	"github.com/bundgaard/battlenet/http"
)

func New(client *http.Client) *Community {
	return &Community{client: client}
}

type Community struct {
	client *http.Client
}

func (c Community) Login() error {
	return c.client.Login()
}

// D3 ACT API

// /d3/data/act
func (c Community) ActIndex() (*types.ActIndex, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/act")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	actIndex := new(types.ActIndex)
	if err := json.NewDecoder(resp.Body).Decode(&actIndex); err != nil {
		return nil, err
	}
	return actIndex, nil
}

// /d3/data/act/:id
func (c Community) Act(id int) (*types.Act, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/act/%d", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	act := new(types.Act)
	if err := json.NewDecoder(resp.Body).Decode(&act); err != nil {
		return nil, err
	}
	return act, nil
}

// D3 ARTISAN and RECIPE API

// /d3/data/artisan/:slug
func (c Community) Artisan(slug string) (*types.Artisan, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/artisan/%s", slug)
	if err != nil {
		// return nil, err
		return nil, err
	}
	defer resp.Body.Close()

	artisan := new(types.Artisan)
	if err := json.NewDecoder(resp.Body).Decode(&artisan); err != nil {
		return nil, err
	}
	return artisan, nil

}

// /d3/data/artisan/:slug/recipe/:recipe
func (c Community) Recipe(slug, recipeSlug string) (*types.Recipe, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/artisan/%s/recipe/%s", slug, recipeSlug)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	recipe := new(types.Recipe)
	if err := json.NewDecoder(resp.Body).Decode(&recipe); err != nil {
		return nil, err
	}
	return recipe, nil

}

// D3 FOLLOWER API

// /d3/data/follower/:slug
func (c Community) Follower(slug string) (*types.Follower, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/follower/%s", slug)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	follower := new(types.Follower)
	if err := json.NewDecoder(resp.Body).Decode(&follower); err != nil {
		return nil, err
	}
	return follower, nil

}

// D3 CHARACTER CLASS and SKILL

// /d3/data/hero/:class
func (c Community) CharacterClass(class string) (*types.Character, error) {

	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/hero/%s", class)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	character := new(types.Character)
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return nil, err
	}
	return character, nil
}

// /d3/data/hero/:class/skill/:skill
func (c Community) Skill(class, slug string) (*types.ApiSkill, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/hero/%s/skill/%s", class, slug)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	skill := new(types.ApiSkill)
	if err := json.NewDecoder(resp.Body).Decode(&skill); err != nil {
		return nil, err
	}
	return skill, nil
}

// D3 ITEM TYPE API

// /d3/data/item-type
func (c Community) ItemTypeIndex() ([]types.ItemTypeIndex, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/item-type")
	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	index := make([]types.ItemTypeIndex, 0)

	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		return nil, err
	}
	return index, nil
}

// /d3/data/item-type/:item
func (c Community) ItemType(item string) (*types.ItemTypeIndex, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/item-type/%s", item)
	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()
	index := new(types.ItemTypeIndex)

	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		return nil, err
	}
	return index, nil
}

// D3 ITEM API

// /d3/data/item/:item-and-id
func (c Community) Item(path string) (*types.ApiItem, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/data/%s", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	item := new(types.ApiItem)
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return nil, err
	}
	return item, nil
}

// D3 PROFILE API

// /d3/profile/:account
func (c Community) Account(account string) (*types.Profile, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/profile/%s/", http.Normalize(account))
	if err != nil {
		return nil, err
	}
	var p types.Profile
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

// /d3/profile/:account/hero/:hero
func (c Community) Hero(account string, hero int) (*types.Hero, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/profile/%s/hero/%d", http.Normalize(account), hero)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var v types.Hero
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

// /d3/profile/:account/hero/:hero/items
func (c Community) DetailedHeroItems(account string, hero int) (interface{}, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/profile/%s/hero/%d/items", account, hero)
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

// /d3/profile/:account/hero/:hero/follower-items
func (c Community) DetailedFollowerItems(account string, hero int) (interface{}, error) {
	resp, err := c.client.TokenRequest("GET", nil, "/d3/profile/%s/hero/%d/follower-items", account, hero)
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
