package models

type JenisGejala struct {
	IDJenisGejala    string         `gorm:"primary_key" json:"id_jenis_gejala"`
	JenisGejala      string         `gorm:"-" json:"jenis_gejala"`
}

func (JenisGejala) TableName() string {
	return "jenis_gejala"
}
