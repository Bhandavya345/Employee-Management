package models

type Employee struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"employee_id"`
	Name       string  `gorm:"size:100;not null" json:"name"`
	Age        int     `gorm:"not null" json:"age"`
	Department string  `gorm:"size:100;not null" json:"department"`
	Salary     float64 `gorm:"type:numeric(10,2);not null" json:"salary"`
	Experience int     `gorm:"not null" json:"experience"`
}
