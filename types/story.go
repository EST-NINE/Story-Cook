package types

type GenerateStoryReq struct {
	Mood     string `json:"mood"`
	Keywords string `json:"keywords"`
}

type CreateStoryReq struct {
	Title    string `json:"title"`
	Mood     string `json:"mood"`
	Keywords string `json:"keywords"`
	Content  string `json:"content"`
}
