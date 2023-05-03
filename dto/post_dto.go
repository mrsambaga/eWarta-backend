package dto

type PostHighlight struct {
	PostId      uint64 `json:"post_id"`
	Title       string `json:"title"`
	SummaryDesc string `json:"summary_desc"`
	ImgUrl      string `json:"img_url"`
	Author      string `json:"author"`
}

type PostDetail struct {
	Title       string `json:"title"`
	SummaryDesc string `json:"summary_desc"`
	ImgUrl      string `json:"img_url"`
	Content     string `json:"content"`
	Author      string `json:"author"`
}
