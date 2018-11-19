package data

import (
	"github.com/jinzhu/gorm"
)

type AIRequest struct {
	gorm.Model
	Text    string
	Intents []Intent
}

type Intent struct {
	Value       string
	AIRequest   *AIRequest
	AIRequestID uint `gorm:"index"`
}