package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"size:100;not null" json:"name"`
	Email    string `gorm:"size:100;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	RoleID   uint   `gorm:"not null;default:2" json:"role_id"`
	Role     string `gorm:"size:20;default:user" json:"role"`
}
