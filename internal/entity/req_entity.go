package entity

type AddContentReq struct {
	Title       string `json:"title"`
	HeadPhoto   string `json:"head_photo"`
	AuthorID    string `json:"AuthorID"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Tag         int64  `json:"tag"`
}

type GetContentListReq struct {
	Tag    int64 `form:"tag"`
	Status int64 `form:"status"`
	Start  int64 `form:"start" binding:"required"`
	Stop   int64 `form:"stop" binding:"required"`
}
