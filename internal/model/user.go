package model

type UserModel struct {
	UserID uint   `gorm:"primaryKey;autoIncrement"`
	Phone  string `json:"phone" validate:"require"`
	Role   string `json:"role" gorm:"not null;default:'user'"`
}
