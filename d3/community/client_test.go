// Copyright (C) 2021 David Bundgaard
package community

import (
	"testing"

	"github.com/bundgaard/battlenet/http"
)

func TestCommunityAPI(t *testing.T) {
	cl := http.New("YOURTOKENHERE", "YOURTOKENHERE", http.Europe)
	c := New(cl)

	if err := c.Login(); err != nil {
		t.Error(err)
	}

	account := "YOURACCOUNTHERE"
	p, err := c.Account(account)
	if err != nil {
		t.Error(err)
	}

	c.DetailedHeroItems(account, p.Heroes[0].ID)
	/*
		for _, hero := range p.Heroes {
			heroInfo, err := c.Hero(account, hero.ID)
			if err != nil {
				t.Error(err)
			}


			c.DetailedHeroItems(account, hero.ID)
		}
	*/

	/* ai, err := c.ActIndex()
	if err != nil {
		t.Error(err)
	}

	t.Log(ai) */

	/* artisan, err := c.Artisan("blacksmith")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Artisan %v", artisan)

	t.Logf("%s", strings.Repeat("=", 80))
	recipe, err := c.Recipe("blacksmith", "hellish-staff-of-herding")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Recipe %v", recipe)
	t.Logf("%s", strings.Repeat("=", 80))

	follower, err := c.Follower("templar")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Follower %v", follower) */
	/* 	t.Logf("%s", strings.Repeat("=", 80))
	   	character, err := c.CharacterClass("crusader")
	   	if err != nil {
	   		t.Error(err)
	   	}
	   	t.Logf("Character %v", character)

	   	t.Logf("%s", strings.Repeat("=", 80))
	   	skill, err := c.Skill("crusader", "punish")
	   	if err != nil {
	   		t.Error(err)
	   	}
	   	t.Logf("skill %v", skill) */

	// c.ItemTypeIndex()
	/* 	if _, err := c.ItemType("scoundrelspecial"); err != nil {
		t.Error(err)
	} */

}
