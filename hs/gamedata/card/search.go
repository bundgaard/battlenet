package card

import (
	"fmt"
	"net/url"

	"github.com/bundgaard/battlenet/internal/ints"
)

func Locale(v string) Option {
	return func(sr *SearchRequest) {
		sr.Locale = v
		sr.values.Add("locale", v)
	}
}
func Set(v string) Option {
	return func(sr *SearchRequest) {
		sr.Set = v
		sr.values.Add("set", v)
	}
}
func Class(v string) Option {
	return func(sr *SearchRequest) {
		sr.Class = v
		sr.values.Add("class", v)
	}
}

func Manacost(is ...int) Option {
	return func(sr *SearchRequest) {
		sr.ManaCost = is
		sr.values.Add("manacost", ints.Join(is, ","))
	}
}

func Attack(is ...int) Option {
	return func(sr *SearchRequest) {
		sr.Attack = is
		sr.values.Add("attack", ints.Join(is, ","))
	}
}

func Health(is ...int) Option {
	return func(sr *SearchRequest) {
		sr.Health = is
		sr.values.Add("health", ints.Join(is, ","))
	}
}

func Collectible(is ...int) Option {
	return func(sr *SearchRequest) {
		sr.Collectible = is
		sr.values.Add("collectible", ints.Join(is, ","))
	}
}

func Rarity(v string) Option {
	return func(sr *SearchRequest) {
		sr.Rarity = v
		sr.values.Add("rarity", v)
	}
}

func Type(v string) Option {
	return func(sr *SearchRequest) {
		sr.Type = v
		sr.values.Add("type", v)
	}
}
func MinionType(v string) Option {
	return func(sr *SearchRequest) {
		sr.MinionType = v
		sr.values.Add("minionType", v)
	}
}
func Keyword(v string) Option {
	return func(sr *SearchRequest) {
		sr.Keyword = v
		sr.values.Add("keyword", v)
	}
}

func TextFilter(v string) Option {
	return func(sr *SearchRequest) {
		sr.TextFilter = v
		sr.values.Add("textFilter", v)
	}
}

func GameMode(v string) Option {
	return func(sr *SearchRequest) {
		sr.GameMode = v
		sr.values.Add("gameMode", v)
	}
}

func Page(n int) Option {
	return func(sr *SearchRequest) {
		sr.Page = n
		sr.values.Add("page", fmt.Sprint(n))
	}
}

func PageSize(v int) Option {
	return func(sr *SearchRequest) {
		sr.PageSize = v
		sr.values.Add("pageSize", fmt.Sprint(v))
	}
}

type SortOrder string

const (
	ManacostAsc      SortOrder = "manaCost:asc"
	ManacostDesc     SortOrder = "manaCost:desc"
	AttackAsc        SortOrder = "attack:asc"
	AttackDesc       SortOrder = "attack:desc"
	HealthAsc        SortOrder = "health:asc"
	HealthDesc       SortOrder = "health:desc"
	ClassAsc         SortOrder = "class:asc"
	ClassDesc        SortOrder = "class:desc"
	GroupByClassAsc  SortOrder = "groupByClass:asc"
	GroupByClassDesc SortOrder = "groupByClass:desc"
	NameAsc          SortOrder = "name:asc"
	NameDesc         SortOrder = "name:desc"
)

func Sort(s SortOrder) Option {
	return func(sr *SearchRequest) {
		sr.Sort = s
		sr.values.Add("sort", string(s))
	}
}

type SearchRequest struct {
	Locale      string // Locality of text, if not set, all available is presented
	Set         string // Slug of the set the cards belong to. If not set all cards will be presented
	Class       string // Slug of the card set
	ManaCost    []int  // Manacost of the card, convert it into a comma separated list
	Attack      []int  // Attack of the card, convert it into a comma separated list
	Health      []int  // Health of the card, convert it into a comma separated list
	Collectible []int  // Collectible or uncollectible (Card that summons other cards), convert it into a comma separated list
	Rarity      string
	Type        string
	MinionType  string
	Keyword     string
	TextFilter  string
	GameMode    string
	Page        int
	PageSize    int
	Sort        SortOrder

	values url.Values
}

type Option func(*SearchRequest)

func Search(opts ...Option) *SearchRequest {
	sr := &SearchRequest{values: url.Values{}}
	for _, opt := range opts {
		opt(sr)
	}

	return sr
}

func (sr *SearchRequest) HasLocale() bool {
	return sr.Locale != ""
}

func (sr *SearchRequest) SetLocale(v string) {
	if sr.Locale != v {
		sr.Locale = v
		sr.values.Add("locale", v)
	}
}

func (sr *SearchRequest) Encode() string {
	return sr.values.Encode()
}

/// UTILITI FUNCTIONS
