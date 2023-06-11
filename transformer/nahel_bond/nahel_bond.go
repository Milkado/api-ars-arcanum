package transformer

type NahelBond struct {
	ID        uint    `json:"id"`
	Order     string  `json:"name"`
	Spren     string  `json:"spren"`
	Attribute string  `json:"attribute"`
	Powers    []Power `json:"powers,string" gorm:"many2many:nahel_bond_powers;"`
}

type Power struct {
	ID            uint        `json:"id"`
	Name          string      `json:"name"`
	MagicSystem   MagicSystem `json:"magic_system"`
	MagicSystemID uint        `json:"magic_system_id"`
}

type MagicSystem struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
}
