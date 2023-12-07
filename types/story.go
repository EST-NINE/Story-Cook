package types

type GenerateStoryReq struct {
	Mood     string `json:"mood" binding:"required"`
	Keywords string `json:"keywords"`
}

type CreateStoryReq struct {
	Title    string `json:"title" binding:"required"`
	Mood     string `json:"mood" binding:"required"`
	Keywords string `json:"keywords" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type ListStoryReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type DeleteStoryReq struct {
	Id uint `json:"id" binding:"required"`
}

type ListStoryResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Mood      string `json:"mood"`
	Keywords  string `json:"keywords"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
