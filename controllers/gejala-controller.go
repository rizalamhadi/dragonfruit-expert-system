package controllers

import (
	h "dragon-fruit-expert-system/helpers"
	models "dragon-fruit-expert-system/models/db_models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func NewGejalaController(db *gorm.DB) *GejalaController {
	return &GejalaController{
		Conn: db,
	}
}

type GejalaController struct {
	Conn *gorm.DB
}

func (g *GejalaController) List(w http.ResponseWriter, r *http.Request) {
	var gejala []models.Gejala

	err := g.Conn.Raw("SELECT id_gejala, gejala, (SELECT jenis_gejala FROM jenis_gejala WHERE jenis_gejala.id_jenis_gejala = gejala.id_jenis_gejala) as id_jenis_gejala FROM dragon.gejala").Scan(&gejala)
	if err.Error != nil {
		h.ResError(w, r, 400, "Terjadi kesalahan..")
		return
	}

	h.ResWithoutMeta(w, r, 200, gejala)
}

func (g *GejalaController) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var gejala models.Gejala
	err := g.Conn.Raw("SELECT id_gejala, gejala, (SELECT jenis_gejala FROM jenis_gejala WHERE jenis_gejala.id_jenis_gejala = gejala.id_jenis_gejala) as id_jenis_gejala FROM dragon.gejala WHERE id_gejala = ?", id).Scan(&gejala)
	if err.Error != nil {
		h.ResError(w, r, 400, "Terjadi kesalahan..")
		return
	}

	h.ResWithoutMeta(w, r, 200, gejala)
}

func (g *GejalaController) Add(w http.ResponseWriter, r *http.Request) {

}
