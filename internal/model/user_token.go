package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserToken struct {
	ID     string `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	UserID string `json:user_id`
	User   User   `json:"user" gorm:foreignKey:UserID"`
	Type   string `json:"type"`
	Token  string `json:"token"`
}

func (t *UserToken) BeforeCreate(_ *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return
}

type AdminToken struct {
	ID     string `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	UserID string `json:"user_id"`
	Type   string `json:"type"`
	Token  string `json:"token"`
}

func (t *AdminToken) BeforeCreate(_ *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return
}
