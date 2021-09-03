package types

import "reflect"

// Hero ...
type Hero struct {
	Alive                    bool             `json:"alive,omitempty"`
	Class                    string           `json:"class"`
	ClassSlug                string           `json:"classSlug"`
	Dead                     bool             `json:"dead"`
	Followers                Followers        `json:"followers,omitempty"`
	Gender                   Gender           `json:"gender"`
	Hardcore                 bool             `json:"hardcore"`
	HighestSoloRiftCompleted int              `json:"highestSoloRiftCompleted,omitempty"`
	ID                       int              `json:"id"`
	Items                    map[string]Item  `json:"items,omitempty"`
	Kills                    map[string]int   `json:"kills"`
	LastUpdated              ProfileTime      `json:"last-updated"`
	LegendaryPowers          []LegendaryPower `json:"legendaryPowers,omitempty"`
	Level                    int              `json:"level"`
	Name                     string           `json:"name"`
	ParagonLevel             int              `json:"paragonLevel"`
	Progress                 Progression      `json:"progression,omitempty"`
	Seasonal                 bool             `json:"seasonal"`
	SeasonCreated            int              `json:"seasonCreated,omitempty"`
	Skills                   Skills           `json:"skills,omitempty"`
	Stats                    Stat             `json:"stats,omitempty"`
}

// SkillCollection used in active and passive ...
type SkillCollection struct {
	Rune  Rune  `json:"rune"`
	Skill Skill `json:"skill"`
}

// Skills ...
type Skills struct {
	Active  []SkillCollection `json:"active"`
	Passive []SkillCollection `json:"passive"`
}

// Progression ...
type Progression struct {
	Act1 Act `json:"act1"`
	Act2 Act `json:"act2"`
	Act3 Act `json:"act3"`
	Act4 Act `json:"act4"`
	Act5 Act `json:"act5"`
}

// LegendaryPower ...
type LegendaryPower struct {
	DisplayColor  string `json:"displayColor"`
	Icon          string `json:"icon"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	TooltipParams string `json:"tooltipParams"`
}

// Followers ...
type Followers struct {
	Enchantress Follower `json:"enchantress"`
	Scoundrel   Follower `json:"scoundrel"`
	Templar     Follower `json:"templar"`
}
type Stat struct {
	ArcaneResist      float64 `json:"arcanceResist"`
	Armor             float64 `json:"armor"`
	AttackSpeed       float64 `json:"attackSpeed"`
	BlockAmountMax    float64 `json:"blockAmountMax"`
	BlockAmountMin    float64 `json:"blockAmountMin"`
	BlockChance       float64 `json:"blockChance"`
	ColdResist        float64 `json:"coldResist"`
	CriticalChance    float64 `json:"critChance"`
	Damage            float64 `json:"damage"`
	Dexterity         float64 `json:"dexterity"`
	FireResist        float64 `json:"fireResist"`
	GoldFind          float64 `json:"goldFind"`
	Healing           float64 `json:"healing"`
	Intelligence      float64 `json:"intelligence"`
	Life              float64 `json:"life"`
	LifeOnHit         float64 `json:"lifeOnHit"`
	LifePerKill       float64 `json:"lifePerKill"`
	Lifesteal         float64 `json:"lifeSteal"`
	LightningResist   float64 `json:"lightningResist"`
	PhysicalResist    float64 `json:"physicalResist"`
	PoisonResist      float64 `json:"poisonResist"`
	PrimaryResource   float64 `json:"primaryResource"`
	SecondaryResource float64 `json:"secondaryResource"`
	Strength          float64 `json:"strength"`
	Thorns            float64 `json:"thorns"`
	Toughness         float64 `json:"toughness"`
	Vitality          float64 `json:"vitality"`
}

func (st Stat) Headers() []string {
	B := reflect.TypeOf(st)
	nFields := B.NumField()
	result := make([]string, nFields)
	for i := 0; i < nFields; i++ {
		result = append(result, B.Field(i).Name)
	}

	return result
}

type Number interface {
}

func (st Stat) Values() map[string]Number {
	T := reflect.TypeOf(st)
	V := reflect.ValueOf(st)
	nFields := T.NumField()
	result := make(map[string]Number)
	for i := 0; i < nFields; i++ {
		typeField := T.Field(i)
		valueField := V.Field(i)
		result[typeField.Name] = Number(valueField)
	}
	return result
}

// ItemType ...
type ItemType struct {
	ID        string `json:"id"`
	TwoHanded bool   `json:"twoHanded"`
}

// ItemAttributes ...
type ItemAttributes struct {
	Primary   []string `json:"primary"`
	Secondary []string `json:"secondary"`
}

// HeroItem ...
type HeroItem struct {
	AccountBound           bool           `json:"accountBound"`
	Armor                  float64        `json:"armor"`
	AttacksPerSecond       float64        `json:"attacksPerSecond"`
	Attributes             ItemAttributes `json:"attributes"`
	AttributesHTML         ItemAttributes `json:"attributesHtml"`
	CraftedBy              ItemCraftedBy  `json:"craftedBy"`
	DisplayColor           string         `json:"displayColor"`
	Dye                    ItemDye        `json:"dye"`
	FlavorText             string         `json:"flavorText"`
	Gems                   []ItemGem      `json:"gems"`
	Icon                   string         `json:"icon"`
	ID                     string         `json:"id"`
	IsSeasonRequiredToDrop bool           `json:"isSeasonRequiredToDrop"`
	ItemLevel              int            `json:"itemLevel"`
	MaxDamage              float64        `json:"maxDamage"`
	MinDamage              float64        `json:"minDamage"`
	Name                   string         `json:"name"`
	OpenSockets            int            `json:"openSockets"`
	RequiredLevel          int            `json:"requiredLevel"`
	SeasonRequiredToDrop   int            `json:"seasonRequiredToDrop"`
	Set                    *HeroItem      `json:"set"`
	Slots                  string         `json:"slots"`
	StackSizeMax           int            `json:"stackSizeMax"`
	TooltipParams          string         `json:"tooltipParams"`
	Type                   ItemType       `json:"type"`
	TypeName               string         `json:"typeName"`
}

// ItemReagent ...
type ItemReagent struct {
	Item     ItemObject `json:"item"`
	Quantity int        `json:"quantity"`
}

// ItemProduced ...
type ItemProduced struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

// ItemCraftedBy ...
type ItemCraftedBy struct {
	Cost         int           `json:"cost"`
	ID           string        `json:"id"`
	ItemProduced ItemProduced  `json:"itemProduced"`
	Name         string        `json:"name"`
	Reagents     []ItemReagent `json:"reagents"`
	Slug         string        `json:"slug"`
}

// ItemObject ...
type ItemObject struct {
	Icon string `json:"icon"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Slug string `json:"slug"`
}

// ItemGem ...
type ItemGem struct {
	Attributes []string   `json:"attributes"`
	IsGem      bool       `json:"isGem"`
	IsJewel    bool       `json:"isJewel"`
	Item       ItemObject `json:"item"`
}

// ItemDye ...
type ItemDye struct {
	Icon          string `json:"icon"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	TooltipParams string `json:"tooltipParams"`
}
