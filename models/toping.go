package models

import "time"

type Toping struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name"`
	Image     string    `json:"toping"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
