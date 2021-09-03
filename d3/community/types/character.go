package types

type Character struct {
	Slug            string              `json:"slug"`
	Name            string              `json:"name"`
	MaleName        string              `json:"maleName"`
	FemaleName      string              `json:"femaleName"`
	Icon            string              `json:"icon"`
	SkillCategories []map[string]string `json:"skillCategories"`
	Skills          map[string][]Skill  `json:"skills"` // can be active or passive

}
