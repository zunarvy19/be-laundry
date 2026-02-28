package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LaundryPackage struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Name             string         `json:"name" gorm:"not null"`
	Price            float64        `json:"price" gorm:"not null"`
	PriceLabel       string         `json:"price_label" gorm:"not null"` // e.g "Rp 15k"
	Unit             string         `json:"unit" gorm:"not null"`        // e.g "/kg"
	SpecialTreatment string         `json:"special_treatment"`           // e.g "Selesai Besok (24 Jam)"
	Benefits         pq.StringArray `json:"benefits" gorm:"type:text[]"` // Array of strings (PostgreSQL specific)
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

type Contact struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	PhoneNumber int  `json:"phone_number" gorm:"not null"`
	IsActive    bool `gorm:"default:true" json:"is_active"`
}

type WebContent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
