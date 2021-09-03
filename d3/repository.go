package d3

import "github.com/bundgaard/battlenet/d3/community/types"

// Repository interface to work against
type Repository interface {
	CreateComparison([]byte) (string, error)
	CreateHero(btag string, hero *types.Hero) error
	CreateHeroItem(btag string, heroID int, heroName string, items map[string]types.HeroItem) error
	CreateItemIcon(string, []byte) error
	CreateItemTooltip(string, string) error
	CreateProfile(string, *types.Profile) error
	FindComparison(id string) ([]byte, error)
	FindHeroesByName(...string) []*types.Hero
	FindHerosByTag(string) ([]*types.Hero, error)
	FindProfileByTag(string) (*types.Profile, error)
	GetHeroDetails(name string) (map[string]types.HeroItem, error)
	GetItemIcon(name string) (*types.Image, error)
	GetItemTooltip(string) (*types.Tooltip, error)
	ListTags() ([]string, error)
	RegisterProfile(userId string, battleTag string) error
}
