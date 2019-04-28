package models

import (
	"github.com/jinzhu/gorm"
)

type Gejala struct {
	IDGejala      string `gorm:"primary_key" json:"id_gejala"`
	IDJenisGejala string `gorm:"-" json:"id_jenis_gejala"`
	Gejala        string `gorm:"-" json:"gejala"`
}

func (Gejala) TableName() string {
	return "gejala"
}

func (g Gejala) GetGejalaByID(db *gorm.DB, IDGejala string) (*Gejala, error) {
	err := db.Where("id_gejala = ?", IDGejala).Find(&g)

	if err.RecordNotFound() {
		return nil, nil
	}

	if err.Error != nil {
		return nil, err.Error
	}

	return &g, nil
}