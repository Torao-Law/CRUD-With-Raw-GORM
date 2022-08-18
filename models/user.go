package models

import "time"

// User model struct
type User struct {
	ID        int               `json:"id" gorm:"primary_key:auto_increment"`
	Name      string            `json:"name" gorm:"type: varchar(255)"`
	Email     string            `json:"email" gorm:"type: varchar(255)"`
	Password  string            `json:"password" gorm:"type: varchar(255)"`
	Profile   ProfileResponse   `json:"profile"`
	Product   []ProductResponse `json:"product"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// menyiapkan proses respons relasi
type UsersProfileResponse struct { // membuat respon untuk relasi
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Profile ProfileResponse `json:"profile"`
}

// fungsi agar memberitahu gorm usersprofileresponse bukan sebuah struck yang akan di migrasi ke database
func (UsersProfileResponse) TableName() string {
	return "users"
}
