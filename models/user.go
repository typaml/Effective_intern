package models

import "time"

type User struct {
	ID             uint   `gorm:"primaryKey"`
	PassportSeries string `gorm:"type:varchar(4);not null"`
	PassportNumber string `gorm:"type:varchar(6);not null"`
	Surname        string `gorm:"type:varchar(100)"`
	Name           string `gorm:"type:varchar(100)"`
	Patronymic     string `gorm:"type:varchar(100)"`
	Address        string `gorm:"type:varchar(255)"`
	Tasks          []Task `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Task struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UnfetchingData struct {
	PassportNumber string `json:"passportNumber"`
}
