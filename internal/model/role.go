package model

type Role struct {
	ID       string `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
}