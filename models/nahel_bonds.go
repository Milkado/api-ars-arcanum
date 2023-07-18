package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type NahelBond struct {
	gorm.Model
	Order     string  `json:"name" validate:"nonzero"`
	Spren     string  `json:"spren" validate:"nonzero"`
	Attribute string  `json:"attribute" validate:"nonzero"`
	Powers    []Power `json:"powers,string" gorm:"many2many:nahel_bond_powers;"`
}

type NahelBondsGet struct {
	ID        uint                `json:"id"`
	Order     string              `json:"name"`
	Spren     string              `json:"spren"`
	Attribute string              `json:"attribute"`
	Powers    []PowerForNahelBond `json:"powers,string" gorm:"many2many:nahel_bond_powers;foreignKey:ID;joinForeignKey:NahelBondID;References:ID;JoinReferences:PowerID"`
}

type PowerForNahelBond struct {
	ID            uint                    `json:"id"`
	Name          string                  `json:"name"`
	Description   string                  `json:"description"`
	MagicSystemID uint                    `json:"magic_system_id"`
	MagicSystem   MagicSystemForNahelBond `json:"magic_system"`
}

type MagicSystemForNahelBond struct {
	ID      uint              `json:"id"`
	Name    string            `json:"name"`
	ShardID uint              `json:"shard_id"`
	Shard   ShardForNahelBond `json:"shard" gorm:"foreignKey:ShardID"`
}

type ShardForNahelBond struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (NahelBondsGet) TableName() string {
	return "nahel_bonds"
}

func (PowerForNahelBond) TableName() string {
	return "powers"
}

func (MagicSystemForNahelBond) TableName() string {
	return "magic_systems"
}

func (ShardForNahelBond) TableName() string {
	return "shards"
}

func ValidadeNahelBond(nahelBond *NahelBond) error {
	if err := validator.Validate(nahelBond); err != nil {
		return err
	}

	return nil
}
