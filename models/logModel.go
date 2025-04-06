package models

import "time"

// Log model for logging to the database
type Log struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
