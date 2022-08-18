package cartdto

type CartRequest struct {
	ID        int `json:"ID" validate:"required"`
	UserID    int `gorm:"type: varchar(255)" json:"userID" validate:"required"`
	ProductID int `gorm:"type: varchar(255)" json:"productID" validate:"required"`
	TopingID  int `gorm:"type: varchar(255)" json:"topingID" validate:"required"`
	Qty       int `gorm:"type: varchar(255)" json:"qty" validate:"required"`
	SubAmount int `gorm:"type: varchar(255)" json:"subAmount" validate:"required"`
}
