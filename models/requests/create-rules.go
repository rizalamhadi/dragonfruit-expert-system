package requests

type CreateRules struct {
	IDRules     string    `json:"id_rules" valid:"required"`
	IDPenyakit  string    `json:"id_penyakit" valid:"required"`
	IDGejala    []string  `json:"id_gejala" valid:"required"`
}
