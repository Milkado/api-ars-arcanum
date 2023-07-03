package models

import "gorm.io/gorm"

type NahelBond struct {
	gorm.Model
	Order     string  `json:"name"`
	Spren     string  `json:"spren"`
	Attribute string  `json:"attribute"`
	Powers    []Power `json:"powers,string" gorm:"many2many:nahel_bond_powers;"`
}

type NahelBondReturn struct {
	ID        uint          `json:"id"`
	Order     string        `json:"name"`
	Spren     string        `json:"spren"`
	Attribute string        `json:"attribute"`
	Powers    []PowerReturn `json:"powers,string" gorm:"many2many:nahel_bond_powers;"`
}