package transformer

type MagicSystem struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Prerequisit string   `json:"prerequisit"`
	ShardID     uint     `json:"shard_id"`
	Shard       Shard    `json:"shard"`
	Powers      []Powers `json:"powers"`
}

type Shard struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Powers struct {
	Name          string `json:"name"`
	MagicSystemID uint   `json:"magic_system_id"`
}
