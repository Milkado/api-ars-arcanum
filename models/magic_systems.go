package models

import "gorm.io/gorm"

type MagicSystem struct {
	gorm.Model
	Name        string  `json:"name"`
	Prerequisit string  `json:"prerequisit"`
	ShardID     uint    `json:"shard_id"`
	Shard       Shard   `json:"shard"`
	Powers      []Power `json:"powers" gorm:"foreignKey:MagicSystemID"`
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
