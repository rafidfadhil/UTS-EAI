package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   string `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	Name string `json:"name"`
}

func (c *Category) BeforeCreate(_ *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return
}

type Product struct {
	ID         string    `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	CategoryID string    `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (p *Product) BeforeCreate(_ *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}
