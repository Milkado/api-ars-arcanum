package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Shard struct {
	gorm.Model
	Name         string        `gorm:"type:text;not null" json:"name" validate:"nonzero"`
	Vessel       string        `gorm:"type:text;not null" json:"vessel" validate:"nonzero"`
	Planet       string        `gorm:"type:text;not null" json:"planet" validate:"nonzero"`
	MagicSystems []MagicSystem `json:"magic_systems"`
}

type ShardGet struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Vessel      string                `json:"vessel"`
	Planet      string                `json:"planet"`
	MagicSystems []MagicSystemForShard `json:"magic_systems" gorm:"foreignKey:ShardID"`
}

type MagicSystemForShard struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Prerequisit string `json:"prerequisit"`
	ShardID     uint   `json:"shard_id"`
}

func (ShardGet) TableName() string {
	return "shards"
}

func (MagicSystemForShard) TableName() string {
	return "magic_systems"
}

func ValidadeShard(shard *Shard) error {
	if err := validator.Validate(shard); err != nil {
		return err
	}

	return nil
}
