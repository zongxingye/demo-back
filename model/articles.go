package model

type Article struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Contents string `json:"contents"`
	Title    string `json:"title"`
	Spot     int    `json:"spot"`
	CreateAt int    `json:"create_at"`
	UpdateAt int    `json:"update_at"`
}
