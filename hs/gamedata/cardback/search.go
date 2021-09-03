package cardback

import "net/url"

type SortOrder string

const (
	NameAsc       SortOrder = "name:asc"
	NameDesc      SortOrder = "name:desc"
	DateAddedAsc  SortOrder = "dateAdded:asc"
	DateAddedDesc SortOrder = "dateAdded:desc"
)

func Locale(v string) Option {
	return func(sr *SearchRequest) {
		sr.Locale = v
		sr.values.Add("locale", v)
	}
}

func Category(v string) Option {
	return func(sr *SearchRequest) {
		sr.Category = v
		sr.values.Add("category", v)
	}
}

func Text(v string) Option {
	return func(sr *SearchRequest) {
		sr.Text = v
		sr.values.Add("textFilter", v)
	}
}

func Sort(s SortOrder) Option {
	return func(sr *SearchRequest) {
		sr.Sort = s
		sr.values.Add("sort", string(s))
	}
}

type SearchRequest struct {
	Locale   string
	Category string
	Text     string
	Sort     SortOrder

	values url.Values
}

type Option func(*SearchRequest)

func Search(opts ...Option) *SearchRequest {
	cbs := &SearchRequest{values: url.Values{}}

	for _, opt := range opts {
		opt(cbs)
	}

	return cbs
}

func (s SearchRequest) Encode() string {
	return s.values.Encode()
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
