package model

// User 用户类
type User struct {
	Id       int64  `json:"id" xorm:"id"`
	UserName string `json:"user_name" xorm:"username"`
	Pwd      string `json:"password" xorm:"password"`
	Nickname string `json:"nick_name" xorm:"nickname"`
	CreateAt int    `json:"create_at" xorm:"create_at"`
	UpdateAt int    `json:"update_at" xorm:"update_at"`
}
