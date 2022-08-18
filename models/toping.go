package models

type Toping struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Price int    `json:"price"`
}

// type TopingResponse struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Price int    `json:"price"`
// 	Image string `json:"image"`
// }

// func (TopingResponse) TableName() string {
// 	return "toping"
// }
