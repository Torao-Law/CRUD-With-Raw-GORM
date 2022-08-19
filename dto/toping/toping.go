package topingdto

type TopingRequest struct {
	Name   string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id" gorm:"type: int"`
}

// type TopingResponse struct {
// 	Name   string `json:"name"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	UserID int    `json:"-"`
// }
