package models

type Mail struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index"`
	Text string `json:"text"`
}
