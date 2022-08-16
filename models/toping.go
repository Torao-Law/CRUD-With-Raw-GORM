package models

import "time"

type Toping struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name      string               `json:"name"`
	Image     string               `json:"image"`
	Price     int                  `json:"price"`
	UserID    int                  `json:"toping_id"`
	User      UsersProfileResponse `json:"user"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type TopingResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
