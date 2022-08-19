package models

type Toping struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Price  int    `json:"price"`
	UserID int    `json:"-"`
}
