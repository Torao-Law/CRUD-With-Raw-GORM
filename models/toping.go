package models

import "time"

type Toping struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"title"`
	Price     int       `json:"price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
