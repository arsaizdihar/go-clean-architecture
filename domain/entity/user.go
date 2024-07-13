package entity

type User struct {
	Base
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"-"`
}
