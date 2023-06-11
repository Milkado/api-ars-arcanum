package transformer

type AllomanticTable struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Physical    string `json:"physical"`
	Enhancement string `json:"enhancement"`
	PowerID     uint   `json:"power_id"`
	Power       Power  `json:"power"`
}

type Power struct {
	ID            uint        `json:"id"`
	Name          string      `json:"name"`
	MagicSystemID uint        `json:"magic_system_id"`
	MagicSystem   MagicSystem `json:"magic_system"`
}

type MagicSystem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
