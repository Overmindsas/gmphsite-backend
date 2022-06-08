package model

type Quote struct {
	Id              int    `json:"id"`
	Dialoge         bool   `json:"dialoge"`
	Private         bool   `json:"private"`
	Tags            string `json:"tags"`
	URL             string `json:"url"`
	FavoritesCount  int    `json:"favorite_count"`
	UpvotesCount    int    `json:"upvotes_count"`
	DownvotesCount  int    `json:"downvotes_count"`
	Author          string `json:"author"`
	AuthorPermalink string `json:"author_permalink"`
	Body            string `json:"body"`
}

type Data struct {
	QotdDate string `json:"qotd_date"`
	Quote    Quote
}

type DataShort struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}
