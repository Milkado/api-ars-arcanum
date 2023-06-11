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