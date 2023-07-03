package models

import "gorm.io/gorm"

type Power struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	MagicSystemID int `json:"magic_system_id"`
	MagicSystem MagicSystem `json:"magic_system"`
}

type PowerReturn struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	MagicSystemID int `json:"magic_system_id"`
	MagicSystem MagicSystemReturn `json:"magic_system" gorm:"foreignKey:MagicSystemID"`
}