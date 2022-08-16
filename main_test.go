package main

import (
	"Eat100/shop_detail_scraping"
	"testing"
)

func BenchmarkScraping(b *testing.B) {
	for i := 0; i < 1; i++ {
		URL := "https://award.tabelog.com/hyakumeiten/izakaya/2021/"
		shop_detail_scraping.ShopDetailScraping(URL)
	}
}
