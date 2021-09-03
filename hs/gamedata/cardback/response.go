package cardback

type CardBack struct {
	ID           int    `json:"id"`
	SortCategory int    `json:"sortCategory"`
	Text         string `json:"text"` // this has the potential to be a map, but as we always default to en_US, it should be fine
	Name         string `json:"name"`
	Image        string `json:"image"`
	Slug         string `json:"slug"`
}

type Response struct {
	CardBacks []CardBack `json:"cardBacks"`
	CardCount int        `json:"cardCount"`
	PageCount int        `json:"pageCount"`
	Page      int        `json:"page"`
}
