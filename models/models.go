package models

type TaskModel struct {
	ID          uint   `gorm:"primaryKey"`
	NumTask     int32  `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Completed   bool   `gorm:"not null"`
}
