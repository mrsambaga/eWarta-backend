package dto

type PostHighlight struct {
	Title       string `json:"title"`
	SummaryDesc string `json:"summary_desc"`
	ImgUrl      string `json:"img_url"`
	Author      string `json:"author"`
}
