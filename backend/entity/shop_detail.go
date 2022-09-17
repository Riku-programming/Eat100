package entity

import (
	"gorm.io/gorm"
	"reflect"
	"strings"
)

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

type SearchResultParams struct {
	Category   string
	Address    string
	Reservable string
	Payment    string
}

func StructToMap(data *SearchResultParams) map[string]string {
	result := make(map[string]string)
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	for i := 0; i < size; i++ {
		field := strings.ToLower(elem.Type().Field(i).Name)
		value := elem.Field(i).String()
		result[field] = value
	}
	return result
}
