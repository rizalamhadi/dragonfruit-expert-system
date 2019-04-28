package models

type RuleDetail struct {
	IDRuleDetail  int     `gorm:"primary_key" json:"id_rule_detail"`
	IDRule        string  `json:"id_rule"`
	IDGejala      string  `json:"id_gejala"`
}

