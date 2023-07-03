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

type MagicSystemReturn struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Prerequisit string `json:"prerequisit"`
	ShardID     uint   `json:"shard_id"`
}

type Tabler interface {
	TableName() string
}

func (MagicSystemReturn) TableName() string {
	return "magic_systems"
}