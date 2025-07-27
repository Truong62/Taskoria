package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"uniqueIndex" json:"email"`
	Password  string `json:"-"`
	FullName  string `json:"full_name"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
