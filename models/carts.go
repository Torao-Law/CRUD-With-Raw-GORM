package models

type Cart struct {
	ID            int             `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int             `json:"product_id"`
	Product       ProductResponse `json:"product"`
	TopingID      []int           `json:"toping_id" form:"toping_id" gorm:"-"`
	Toping        []Toping        `json:"toping" gorm:"many2many:cart_toping"`
	TransactionID int             `json:"transcation_id"`
	Transaction   Transaction     `json:"-"`
	Qty           int             `json:"qty"`
	SubAmount     int             `json:"sub_amount"`
}

type CartResponse struct {
	ID        int             `json:"id"`
	ProductID int             `json:"product_id"`
	ToppingID int             `json:"topping_id"`
	Product   ProductResponse `json:"product"`
	Toping    []Toping        `json:"toping"`
	SubAmount int             `json:"sub_amount"`
}

func (CartResponse) TableName() string {
	return "carts"
}
