package models

import "gorm.io/gorm"

type AllomanticTable struct {
	gorm.Model
	Name string `json:"name"`
	Physical string `json:"physical"`
	Enhancement string `json:"enhancement"`
	PowerID int `json:"power_id"`
	Power Power `json:"power"`
}

type AllomanticTableReturn struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Physical string `json:"physical"`
	Enhancement string `json:"enhancement"`
	PowerID int `json:"power_id"`
	Power PowerReturn `json:"power" gorm:"foreignKey:PowerID"`
}