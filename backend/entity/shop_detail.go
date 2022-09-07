package entity

import "gorm.io/gorm"

type Restaurant struct {
	// fixme gorm.Modelはomitemptyが効かないのでgorm.Modelではなく代わりにIDカラムを追加した方がいいかも
	gorm.Model

	Category    string `json:",omitempty"`
	ShopName    string `json:",omitempty"`
	Reservable  bool   `json:",omitempty"`
	Address     string `json:",omitempty"`
	Time        string `json:",omitempty"`
	Payment     bool   `json:",omitempty"`
	PhoneNumber string `json:",omitempty"`
	Cost        int    `json:",omitempty"`
	URL         string `json:",omitempty"`
}
