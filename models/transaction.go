package models

type Transaction struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID int                  `json:"user_id"`
	User   UsersProfileResponse `json:"user" grom:"-"`
	Carts  []Cart               `json:"cart"`
	Amount int                  `json:"amount"`
}
