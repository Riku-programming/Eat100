package category

import (
	"Eat100/entity"
	"fmt"
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
	return category
}

func ClassifyCategory(category string) string {
	// todo return ""で返している値を構造体にしたりして管理しやすくしたい
	switch category {
	case entity.Unagi:
		fmt.Println("うなぎ")
		return "うなぎ"
	case entity.Izakaya:
		return "居酒屋"
	default:
		return "カテゴリーなし"
	}
}
