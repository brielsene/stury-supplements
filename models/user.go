package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}
