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
	Title string `json:"title" binding:"required"`
}

type UpdateStoryReq struct {
	Title         string `json:"title" binding:"required"`
	UpdateTitle   string `json:"update_title"`
	UpdateContent string `json:"update_content"`
}

type SelectStoryReq struct {
	Mood string `json:"mood" binding:"required"`
}

type StoryResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Mood      string `json:"mood"`
	Keywords  string `json:"keywords"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
