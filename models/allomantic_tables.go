package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type AllomanticTable struct {
	gorm.Model
	Name        string `json:"name" validate:"nonzero"`
	Physical    string `json:"physical" validate:"nonzero"`
	Enhancement string `json:"enhancement" validate:"nonzero"`
	PowerID     int    `json:"power_id" validate:"nonzero"`
	Power       Power  `json:"power"`
}

type AllomanticTableGet struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Physical    string `json:"physical"`
	Enhancement string `json:"enhancement"`
	PowerID     int    `json:"power_id"`
}

type PowerForAllomanticTable struct {
	ID            uint                          `json:"id"`
	Name          string                        `json:"name"`
	MagicSystemID int                           `json:"magic_system_id"`
	MagicSystem   MagicSystemForAllomanticTable `json:"magic_system"`
}

type MagicSystemForAllomanticTable struct {
	ID      uint                    `json:"id"`
	Name    string                  `json:"name"`
	ShardID uint                    `json:"shard_id"`
	Shard   ShardForAllomanticTable `json:"shard"`
}

type ShardForAllomanticTable struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (AllomanticTableGet) TableName() string {
	return "allomantic_tables"
}

func (PowerForAllomanticTable) TableName() string {
	return "powers"
}

func (MagicSystemForAllomanticTable) TableName() string {
	return "magic_systems"
}

func (ShardForAllomanticTable) TableName() string {
	return "shards"
}

func ValidadeAllomanticTable(allomanticTable *AllomanticTable) error {
	if err := validator.Validate(allomanticTable); err != nil {
		return err
	}

	return nil
}
