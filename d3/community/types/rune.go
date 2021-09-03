// Copyright (C) 2021 David Bundgaard
package types

type ApiSkill struct {
	Skill Skill  `json:"skill"`
	Runes []Rune `json:"runes"`
}
type Rune struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	HTML        string `json:"descriptionHtml"`
}
