package view

//  articleテーブルデータ
type Article struct {
	Article_id string ` json:"article_id" `
	Image_path string `json:"image_path" `
	Title string `json:"title"`
	Context string `json:"context" `
	Genre   int `json:"Genre" `
	Nice    int `json:"Nice" `
	EraYear int  `json:"EraYear"`
	EraMonth int  `json:"EraMonth"`
}