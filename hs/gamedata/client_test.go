// Copyright (C) 2021 David Bundgaard
package gamedata

import (
	"testing"

	"github.com/bundgaard/battlenet/hs/gamedata/card"
	"github.com/bundgaard/battlenet/hs/gamedata/cardback"
	"github.com/bundgaard/battlenet/http"
)

func TestGameData(t *testing.T) {
	httpClient := http.New("YOUROWNTOKENEHRE", "YOUROWNTOKENEHRE", http.Europe)
	c := NewWith(httpClient)

	if err := c.Login(); err != nil {
		t.Error(err)
	}

	c.CardSearch(card.Search(card.Set("wild")))

	c.Metadata()

	_, err := c.CardBackSearch(cardback.Search(cardback.Category("base")))
	if err != nil {
		t.Error(err)
	}

	if _, err := c.CardBackBy("25", ""); err != nil {
		t.Error(err)
	}
}
