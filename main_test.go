package main

import "testing"

func BenchmarkScraping(b *testing.B) {
	for i := 0; i < 1; i++ {
		URL := "https://award.tabelog.com/hyakumeiten/izakaya/2021/"
		scraping(URL)
	}
}
