package transformer

type Shard struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	Planet       string        `json:"planet"`
	Vessel       string        `json:"vessel"`
	MagicSystems []MagicSystem `json:"magic_systems"`
}

type MagicSystem struct {
	Name        string `json:"name"`
	Prerequisit string `json:"prerequisit"`
	ShardID     uint   `json:"shard_id"`
}