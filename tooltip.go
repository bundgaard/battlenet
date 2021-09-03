// Copyright (C) 2021 David Bundgaard
package battlenet

type Tooltip struct {
	ID                     string                   `json:"id"`
	Slug                   string                   `json:"slug"`
	Name                   string                   `json:"name"`
	Icon                   string                   `json:"icon"`
	TooltipParams          string                   `json:"tooltipParams"`
	RequiredLevel          int                      `json:"requiredLevel"`
	Accountbound           bool                     `json:"accountBound"`
	Flavortext             string                   `json:"flavorText"`
	FlavorHTML             string                   `json:"flavorTextHtml"`
	TypeName               string                   `json:"typeName"`
	Type                   TooltipType              `json:"type"`
	Color                  string                   `json:"color"`
	IsSeasonRequiredToDrop bool                     `json:"isSeasonRequiredToDrop"`
	Slots                  []string                 `json:"slots"`
	Attributes             map[string]interface{}   `json:"attributes"`
	RandomAffixes          []map[string]interface{} `json:"randomAffixes"`
	SetName                string                   `json:"setName"`
	SetNameHTML            string                   `json:"setNameHtml"`
	SetDescription         string                   `json:"setDescription"`
	SetDescriptionHTML     string                   `json:"setDescriptionHtml"`
	SetItems               []string                 `json:"setItems"`
}

type TooltipType struct {
	Twohandded bool   `json:"twoHanded"`
	ID         string `json:"id"`
}
