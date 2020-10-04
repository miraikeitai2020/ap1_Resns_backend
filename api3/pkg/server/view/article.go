package view

//
type ArticleRequest struct {
	Genre string `json:genre`
	Month string `json:month`
	Year string `json:year`
}
type ArticleDetailRequest struct {
	ArticleID string `json:articleID`
}
//  記事テーブルデータ
type Article struct {
	Title string `json:"title"`
	ImagePath string `json:"imagePath" `
	Context string `json:"context" `
	Nice string `json:nice`
	Comment []Comment `json:comments`
}
//  コメントテーブルデータ
type Comment struct {
	UserName string `json:"userName"`
	UserImage string `json:"userImage"`
	Contents string `json:"contents"`
}