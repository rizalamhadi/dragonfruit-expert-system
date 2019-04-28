package models

type Rule struct {
	IDRules     string  `gorm:"primary_key" json:"id_rules"`
	IDPenyakit  string  `json:"id_penyakit"`
	IDGejala    string  `gorm:"-" json:"id_gejala"`
}
