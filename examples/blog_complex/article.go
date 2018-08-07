package main

// The Article struct holds some information about a blog article
type Article struct {
	ArticleID string
	AuthorID  string
	Text      string
}

// Articles returns all the articles in the blog
func Articles() []Article {
	return []Article{
		{
			ArticleID: "a1",
			AuthorID:  "u1",
			Text:      "Five-star WR Blake Miller signs letter of intent...",
		},
		{
			ArticleID: "a2",
			AuthorID:  "u2",
			Text:      "Late in the fourth quarter, senior quarterback Riley...",
		},
		{
			ArticleID: "a3",
			AuthorID:  "u3",
			Text:      "If last week's scrimmage is any indicator, this season...",
		},
	}
}
