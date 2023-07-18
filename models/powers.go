package models

import "gorm.io/gorm"

type Power struct {
	gorm.Model
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	MagicSystemID int         `json:"magic_system_id"`
	MagicSystem   MagicSystem `json:"magic_system"`
}

type PowerGet struct {
	ID            uint                `json:"id"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	MagicSystemID int                 `json:"magic_system_id"`
	MagicSystem   MagicSystemForPower `json:"magic_system"`
}

type MagicSystemForPower struct {
	ID      uint          `json:"id"`
	Name    string        `json:"name"`
	ShardID uint          `json:"shard_id"`
	Shard   ShardForPower `json:"shard" gorm:"foreignKey:ShardID"`
}

type ShardForPower struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (PowerGet) TableName() string {
	return "powers"
}

func (MagicSystemForPower) TableName() string {
	return "magic_systems"
}

func (ShardForPower) TableName() string {
	return "shards"
}
