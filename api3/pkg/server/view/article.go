package view

//  articleテーブルデータ
type Article struct {
	Article_id int ` json:"article_id" `
	Image_path string `json:"image_path" `
	Title string `json:"title"`
	Context string `json:"context" `
	Genre    int `json:"genre" `
	Tag string  `json:"tag"`
}