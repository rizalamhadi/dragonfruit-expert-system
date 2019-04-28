package controllers

import (
	h "dragon-fruit-expert-system/helpers"
	models "dragon-fruit-expert-system/models/db_models"
	"dragon-fruit-expert-system/models/requests"
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func NewRuleController(db *gorm.DB) *RuleController {
	return &RuleController{
		Conn: db,
	}
}

type RuleController struct {
	Conn *gorm.DB
}

func (p *RuleController) List(w http.ResponseWriter, r *http.Request) {
	var rules []models.Rule

	err := p.Conn.Raw("SELECT id_rules, (SELECT CONCAT('[', penyakit.id_penyakit, '] ', penyakit.nama_penyakit) FROM penyakit WHERE rules.id_penyakit = penyakit.id_penyakit) as id_penyakit, (SELECT GROUP_CONCAT(id_gejala, '') as id_gejala FROM rule_details WHERE rules.id_rules = rule_details.id_rule GROUP BY rule_details.id_rule) as id_gejala FROM dragon.rules").Scan(&rules)

	if err.Error != nil {
		h.ResError(w, r, 400, "Terjadi kesalahan..")
		return
	}

	h.ResWithoutMeta(w, r, 200, rules)
}

func (p *RuleController) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var rule models.Rule

	err := p.Conn.Where("id_rules = ?", id).Find(&rule)
	if err.Error != nil {
		h.ResError(w, r, 400, "Terjadi kesalahan..")
		return
	}

	h.ResWithoutMeta(w, r, 200, rule)
}

func (p *RuleController) Create(w http.ResponseWriter, r *http.Request) {
	var param requests.CreateRules
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		h.ResError(w, r, 200, "Kolom belum terisi, lengkapi terlebih dahulu")
		return
	}

	_, errors := govalidator.ValidateStruct(param)
	if errors != nil {
		h.ResError(w, r, 200, "Kolom belum terisi, lengkapi terlebih dahulu")
		return
	}

	var checkKode models.Rule
	err := p.Conn.Where("id_rules = ?", param.IDRules).First(&checkKode)
	if !err.RecordNotFound() {
		h.ResError(w, r, 200, "Kode sudah digunakan")
		return
	}

	rules := models.Rule{
		IDRules:    param.IDRules,
		IDPenyakit: param.IDPenyakit,
	}
	res := p.Conn.Model(&models.Rule{}).Create(&rules)
	if res.Error != nil {
		h.ResError(w, r, 200, "Cannot create !")
		return
	}

	for i := 0; i < len(param.IDGejala); i++ {
		ruleDetail := models.RuleDetail{
			IDRule:   param.IDRules,
			IDGejala: param.IDGejala[i],
		}
		res := p.Conn.Model(&models.RuleDetail{}).Create(&ruleDetail)
		if res.Error != nil {
			h.ResError(w, r, 200, "Cannot create !")
			return
		}
	}

	h.ResWithoutMeta(w, r, 201, rules)
}
