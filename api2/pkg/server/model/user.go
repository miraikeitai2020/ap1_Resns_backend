package model

//  userテーブルデータ
type User struct {
	Id     string `json:"id" `
	Name   string `json:"name" `
	Image  string `json:"image" `
	Year   int    `json:"year" `
	month  int    `json:"month" `
	Gender int    `json:"gender" `
}

type Retoken struct {
	Token string `json:"token" `
}
