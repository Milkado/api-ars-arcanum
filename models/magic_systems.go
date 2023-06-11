package models

import "gorm.io/gorm"

type MagicSystem struct {
	gorm.Model
	Name        string `json:"name"`
	Prerequisit string `json:"prerequisit"`
	ShardID     uint   `json:"shard_id"`
	Shard       Shard  `json:"shard"`
	Powers      []Power `json:"powers"`
}
