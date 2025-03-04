package models

type User struct {
	ID            int    `json:"id" gorm:"primary_key"`
	Username      string `json:"username" gorm:"unique"`
	Email         string `json:"email" gorm:"unique"`
	EmailVerified bool   `json:"email_verified" gorm:"default:false"`
}
