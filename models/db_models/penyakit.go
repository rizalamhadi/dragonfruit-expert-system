package models

type Penyakit struct {
	IDPenyakit       string         `gorm:"primary_key" json:"id_penyakit"`
	NamaPenyakit     string         `gorm:"-" json:"nama_penyakit"`
	SolusiPenyakit   string         `gorm:"-" json:"solusi_penyakit"`
}

func (Penyakit) TableName() string {
	return "penyakit"
}
