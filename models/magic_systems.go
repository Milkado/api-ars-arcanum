package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type MagicSystem struct {
	gorm.Model
	Name        string  `json:"name" validate:"nonzero"`
	Prerequisit string  `json:"prerequisit" validate:"nonzero"`
	ShardID     uint    `json:"shard_id" validate:"nonzero"`
	Shard       Shard   `json:"shard"`
	Powers      []Power `json:"powers" gorm:"foreignKey:MagicSystemID"`
}

type MagicSystemPost struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" validate:"nonzero"`
	Prerequisit string `json:"prerequisit" validate:"nonzero"`
	ShardID    uint   `json:"shard_id" validate:"nonzero"`
}

type MagicSystemGet struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	Prerequisit string              `json:"prerequisit"`
	ShardID     uint                `json:"shard_id"`
	Shard       ShardForMagicSystem `json:"shard" gorm:"foreignKey:ShardID"`
}

type ShardForMagicSystem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (MagicSystemGet) TableName() string {
	return "magic_systems"
}

func (ShardForMagicSystem) TableName() string {
	return "shards"
}

func ValidadeMagicSystem(magicSystem *MagicSystemPost) error {
	if err := validator.Validate(magicSystem); err != nil {
		return err
	}

	return nil
}
