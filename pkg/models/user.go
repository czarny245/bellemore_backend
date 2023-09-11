package models

type User struct {
  ID            uint    `gorm:"primaryKey"`
  Email         string  `gorm:"uniqueIndex;not null"` 
  PasswordHash  string  `gorm:"not null"`
}
