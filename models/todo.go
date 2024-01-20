package models

type Todo struct {
	ID        uint `gorm:"primaryKey"`
	Task      string
	Completed int 
}
