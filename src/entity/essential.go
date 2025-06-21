package entity

import "time"

type EssentialEntity struct {
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
	Id          int       `gorm:"primaryKey" json:"id"`
}
