package types

type GenerateStoryReq struct {
	Mood     string `json:"mood" binding:"required" example:"开心"`
	Keywords string `json:"keywords" example:"室友+电脑"`
}

type CreateStoryReq struct {
	Title    string `json:"title" binding:"required"  example:"story1"`
	Mood     string `json:"mood" binding:"required" example:"开心"`
	Keywords string `json:"keywords" binding:"required" example:"室友+电脑"`
	Content  string `json:"content" binding:"required" example:"content1"`
}

type ListStoryReq struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}

type DeleteStoryReq struct {
	Title string `json:"title" binding:"required" example:"story1"`
}

type UpdateStoryReq struct {
	Title         string `json:"title" binding:"required" example:"story1"`
	UpdateTitle   string `json:"update_title" example:"story2"`
	UpdateContent string `json:"update_content" example:"content2"`
}

type SelectStoryReq struct {
	Mood string `json:"mood" binding:"required" example:"开心"`
}

type StoryResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Mood      string `json:"mood"`
	Keywords  string `json:"keywords"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
