// Copyright (C) 2021 David Bundgaard
package types

type Artisan struct {
	Slug     string   `json:"slug"`
	Name     string   `json:"name"`
	Portrait string   `json:"portrait"`
	Training Training `json:"training"`
}

type Training struct {
	Tiers []Tier `json:"tiers"`
}

type Tier struct {
	Tier           int      `json:"tier"`
	TrainedRecipes []Recipe `json:"trainedRecipes"`
	TaughtRecipes  []Recipe `json:"taughtRecipes"`
}

type Recipe struct {
	ID           string    `json:"id"`
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	Cost         int       `json:"cost"`
	Reagents     []Reagent `json:"reagents"`
	ItemProduced Item      `json:"itemProduced"`
}

type Reagent struct {
	Quantity int  `json:"quantity"`
	Item     Item `json:"item"`
}

type Item struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

type ApiItem struct {
	Item

	Tooltip       string `json:"tooltipParams"`
	RequiredLevel int    `json:"requiredLevel"`
	StackSizeMax  int    `json:"stackSizeMax"`
	AccountBound  bool   `json:"accountBound"`
	TypeName      string `json:"typeName"`
	Type          struct {
		Twohanded bool   `json:"twoHanded"`
		ID        string `json:"id"`
	} `json:"type"`
	Color                  string        `json:"color"`
	Description            string        `json:"description"`
	HTML                   string        `json:"descriptionHtml"`
	IsSeasonRequiredToDrop bool          `json:"isSeasonRequiredToDrop"`
	SeasonRequiredToDrop   int           `json:"seasonRequiredToDrop"`
	Slots                  []interface{} `json:"slots"`
	SetItems               []interface{} `json:"setItems"`
}
