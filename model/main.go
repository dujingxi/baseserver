package model

import "time"

type Service struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	NicknameZh string    `json:"nicknameZH"`
	NicknameEn string    `json:"nicknameEN"`
	Avswitch   string    `json:"avswitch"`
	Desc       string    `json:"desc"`
	CreateTime time.Time `json:"datetime"`
}
