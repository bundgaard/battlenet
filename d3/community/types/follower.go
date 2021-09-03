package types

type Follower struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	RealName string `json:"realName"`
	Portrait string `json:"templar"`

	Skills []Skill `json:"skills"`
}

type Skill struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Level       int    `json:"level"`
	TooltipURL  string `json:"tooltipUrl"`
	Description string `json:"description"`
	HTML        string `json:"descriptionHtml"`
	Flavor      string `json:"flavorText,omitempty"`
}
