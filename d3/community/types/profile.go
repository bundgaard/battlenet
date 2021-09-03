// Copyright (C) 2021 David Bundgaard
package types

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

// Profile ...
type Profile struct {
	BattleTag            string                     `json:"battleTag"`
	ParagonLevel         int                        `json:"paragonLevel"`
	ParagonLevelSeason   int                        `json:"paragonLevelSeason"`
	GuildName            string                     `json:"guildName"`
	Heroes               []ProfileHero              `json:"heroes"`
	LastHeroPlayed       int                        `json:"lastHeroPlayed"`
	LastUpdated          ProfileTime                `json:"lastUpdated"`
	Kills                map[string]int             `json:"kills"`
	HighestHardcoreLevel int                        `json:"highestHardcoreLevel"`
	TimePlayed           map[string]float64         `json:"timePlayed"` // should fix to a time format for better visibility
	Progression          map[string]bool            `json:"progression"`
	SeasonalProfiles     map[string]SeasonalProfile `json:"seasonalProfiles"`
	Blacksmith           Artisan                    `json:"blacksmith"`
	Jeweler              Artisan                    `json:"jeweler"`
	Mystic               Artisan                    `json:"mystic"`
	BlacksmithSeason     Artisan                    `json:"blacksmithSeason"`
	MysticSeason         Artisan                    `json:"mysticSeason"`
	JewelerSeason        Artisan                    `json:"jewelerSeason"`
}

// ProfileHero ...
type ProfileHero struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Class        string         `json:"class"`
	ClassSlug    string         `json:"classSlug"`
	Gender       int            `json:"gender"`
	Level        int            `json:"level"`
	Kills        map[string]int `json:"kills"`
	ParagonLevel int            `json:"paragonLevel"`
	Hardcore     bool           `json:"hardcore"`
	Seasonal     bool           `json:"seasonal"`
	Dead         bool           `json:"dead"`
	LastUpdated  ProfileTime    `json:"last-updated"`
}

// SeasonalProfile ...
type SeasonalProfile struct {
	SeasonID             int                `json:"seasonId"`
	ParagonLevel         int                `json:"paragonLevel"`
	ParagonLevelHardcore int                `json:"paragonLevelHardcore"`
	Kills                map[string]int     `json:"kills"`
	TimePlayed           map[string]float64 `json:"timePlayed"`
	HighestHardcoreLevel int                `json:"highestHardcoreLevel"`
}

// Gender ...
type Gender string

// UnmarshalJSON ...
func (g *Gender) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		log.Println("failed to parse gender")
		return err
	}

	switch v {
	case 0:
		*g = "Male"
	case 1:
		*g = "Female"
	default:
		*g = "Genderless"
	}
	return nil
}

// MarshalJSON ...
func (g Gender) MarshalJSON() ([]byte, error) {
	var s int
	switch strings.ToLower(string(g)) {
	case "male":
		s = 0
	case "female":
		s = 1
	default:
		s = -1
	}
	return json.Marshal(s)
}
