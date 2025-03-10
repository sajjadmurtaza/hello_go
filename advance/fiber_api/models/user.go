package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
}
