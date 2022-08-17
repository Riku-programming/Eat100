package category

import (
	"Eat100/entity"
	"log"
	"net/url"
	"strings"
)

func ExtractCategoryFromURL(URL string) string {
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	// URLが "https://award.tabelog.com/hyakumeiten/unagi"の場合unagiが返され
	// URLが "https://award.tabelog.com/hyakumeiten/izakaya/2021"の場合izakayaが返される
	category := strings.Split(u.Path, "/")[2]
	// カテゴリによってはURLが違う。
	//例えば居酒屋などは/izakayaだが寿司などは/sushi_tokyoのようにエリア分けされている
	if strings.Contains(category, "_") {
		category := strings.Split(category, "_")[0]
		return category
	}
	return category
}

func ClassifyCategory(category string) string {
	// todo return ""で返している値を構造体にしたりして管理しやすくしたい
	switch category {
	case entity.Unagi:
		return "うなぎ"
	case entity.Izakaya:
		return "居酒屋"
	case entity.Sushi:
		return "寿司"
	default:
		return "カテゴリーなし"
	}
}
