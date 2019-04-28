package controllers

import (
	h "dragon-fruit-expert-system/helpers"
	models "dragon-fruit-expert-system/models/db_models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func NewPenyakitController(db *gorm.DB) *PenyakitController {
	return &PenyakitController{
		Conn: db,
	}
}

type PenyakitController struct {
	Conn *gorm.DB
}

func (p *PenyakitController) List(w http.ResponseWriter, r *http.Request) {
	var penyakit []models.Penyakit

	err := p.Conn.Find(&penyakit)
	if err.Error != nil {
		h.ResError(w, r, 400, "Terjadi kesalahan..")
		return
	}

	h.ResWithoutMeta(w, r, 200, penyakit)
}

func (p *PenyakitController) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var penyakit models.Penyakit

	err := p.Conn.Where("id_penyakit = ?", id).Find(&penyakit)
	if err.Error != nil {
		h.ResError(w, r, 400, "Terjadi kesalahan..")
		return
	}

	h.ResWithoutMeta(w, r, 200, penyakit)
}
