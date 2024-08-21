package models

import "time"

type Posts struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(200)"`
	Content     string    `json:"content" gorm:"type:text"`
	Category    string    `json:"category" gorm:"type:varchar(100)"`
	CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"`
	Status      string    `json:"status" gorm:"type:varchar(100)"`
}
