package model

type User struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email" grom:"unique;not null"`
	Password   []byte `json:"-"`
	IsVerified bool   `gorm:"default:false"`
}
