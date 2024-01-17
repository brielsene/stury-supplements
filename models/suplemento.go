package models

import "gorm.io/gorm"

type Suplementos struct {
	gorm.Model
	Nome       string  `json:"nome"`
	Historia   string  `json:"historia"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
}
