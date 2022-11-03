package entity

type BaseContent struct {
	Title     string `json:"title"`
	HeadPhoto string `json:"head_photo"`
	ContentID string `json:"content_id"`
	Name      string `json:"name"`
	AuthorID  string `json:"author_id"`
	Tag       int64  `json:"tag"`
}

type Content struct {
	Title       string `json:"title"`
	Tag         int64  `json:"tag"`
	HeadPhoto   string `json:"head_photo"`
	ContentID   string `json:"content_id"`
	Name        string `json:"name"`
	AuthorID    string `json:"author_id"`
	Status      int64  `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	Description string `json:"description"`
}
