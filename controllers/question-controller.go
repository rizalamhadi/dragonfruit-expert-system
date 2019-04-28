package controllers

import (
	h "dragon-fruit-expert-system/helpers"
	models "dragon-fruit-expert-system/models/db_models"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/jinzhu/gorm"
)

func NewQuestionController(db *gorm.DB) *QuestionController {
	return &QuestionController{
		Conn: db,
	}
}

type QuestionController struct {
	Conn *gorm.DB
}

type Param struct {
	ParentNode string   `json:"parent_node"`
	OpenNode   string   `json:"open_node"`
	CloseNodes []string `json:"close_nodes"`
	Jawab      bool     `json:"jawab"`
	Rules      []string `json:"rules"`
}

func (q *QuestionController) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var param Param
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		h.ResError(w, r, 400, "Masukkan node terlebih dahulu")
		return
	}

	var closeNodes []string
	if param.OpenNode != "" {
		closeNodes = append(param.CloseNodes, param.OpenNode)
	} else {
		closeNodes = param.CloseNodes
	}

	var rules []string
	if param.OpenNode != "" && param.Jawab {
		rules = append(param.Rules, param.OpenNode)
	} else {
		rules = param.Rules
	}

	nextOpenNode := getRules(q.Conn, param, closeNodes)
	if nextOpenNode == "" {
		IDRules := checkDisease(q.Conn, rules)

		response := make(map[string]interface{})
		response["rules"] = rules
		response["close_nodes"] = closeNodes

		if IDRules != "" {
			var penyakit models.Penyakit
			IDPenyakit := q.Conn.Table("rules").Where("id_rules = ?", IDRules).Select("id_penyakit").Limit(1).SubQuery()
			_ = q.Conn.Where("id_penyakit = ?", IDPenyakit).First(&penyakit)

			response["penyakit"] = penyakit
			h.ResWithoutMeta(w, r, 200, response)
			return
		}

		response["penyakit"] = "Penyakit tidak ditemukan"
		h.ResWithoutMeta(w, r, 200, response)
		return
	}

	var g models.Gejala
	gejala, err := g.GetGejalaByID(q.Conn, nextOpenNode)
	if err != nil {
		h.ResError(w, r, 400, err.Error())
		return
	}

	response := make(map[string]interface{})
	response["open_node"] = gejala.IDGejala
	response["pertanyaan"] = "Apakah " + strings.ToLower(gejala.Gejala) + "?"
	response["rules"] = rules
	response["close_nodes"] = closeNodes

	h.ResWithoutMeta(w, r, 200, response)
}

func getRules(db *gorm.DB, param Param, closeNodes []string) string {
	var r []models.RuleDetail
	_ = db.Group("id_rule").Order("id_rule").Select("id_rule, GROUP_CONCAT(id_gejala, '') as id_gejala").Find(&r)

	gejala := make(map[int][]string)
	for k, v := range r {
		gejala[k] = strings.Split(v.IDGejala, ",")
	}

	fmt.Println(gejala)

	if param.OpenNode == "" {
		return gejala[0][0]
	}

	for k, v := range gejala {
		index := SliceIndex(len(v), func(i int) bool { return v[i] == param.OpenNode })

		if index > -1 && index+1 <= len(v) {
			if !param.Jawab {
				if k > 1 && index > 0 {
					for _, v2 := range gejala {
						found := SliceIndex(len(v2), func(i int) bool { return v2[i] == gejala[k][index-1] })
						if len(v2) > index && found > -1 && gejala[k][index-1] == v2[found] {
							isClosed := SliceIndex(len(closeNodes), func(i int) bool { return closeNodes[i] == v2[index] })
							if isClosed <= 0 {
								return v2[index]
							}
						}
					}
					return ""

				} else {
					opened := true
					n := k + 1
					if n < len(gejala) && len(gejala[n]) > index {
						for opened {
							isClosed := SliceIndex(len(closeNodes), func(i int) bool { return closeNodes[i] == gejala[n][index] })
							if isClosed >= 0 {
								n++
							} else {
								opened = false
							}
						}

						if n > 0 && index > 0 {
							if gejala[n][index-1] == gejala[n-1][index-1] {
								return gejala[n][index]
							}
						} else {
							return gejala[n][index]
						}
					}
				}
			} else {
				if index+1 < len(v) {
					isClosed := sort.Search(len(param.CloseNodes), func(i int) bool { return param.CloseNodes[i] == v[index+1] })
					if len(param.CloseNodes) == 0 || isClosed >= len(param.CloseNodes) {
						return v[index+1]
					}
				}
			}
		}
	}

	return ""
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func checkDisease(db *gorm.DB, rules []string) string {
	var r []models.RuleDetail
	_ = db.Group("id_rule").Order("id_rule").Select("id_rule, GROUP_CONCAT(id_gejala, '') as id_gejala").Find(&r)

	for _, v := range r {
		gejala := strings.Split(v.IDGejala, ",")

		checkRules := testEq(gejala, rules)
		if checkRules {
			return v.IDRule
		}
	}

	return ""
}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
