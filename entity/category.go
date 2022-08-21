package entity

type Category struct {
	searchWord   string
	categoryName string
}

var (
	Curry       = Category{searchWord: "curry", categoryName: "カレー"}
	Unagi       = Category{searchWord: "unagi", categoryName: "うなぎ"}
	Yakitori    = Category{searchWord: "yakitori", categoryName: "焼き鳥"}
	Sushi       = Category{searchWord: "sushi_tokyo", categoryName: "寿司"}
	Tonkatsu    = Category{searchWord: "tonkatsu", categoryName: "とんかつ"}
	Soba        = Category{searchWord: "soba", categoryName: "そば"}
	Hamburger   = Category{searchWord: "hamburger", categoryName: "ハンバーガー"}
	Okonomiyaki = Category{searchWord: "okonomiyaki", categoryName: "お好み焼き"}
	Tempura     = Category{searchWord: "tempura", categoryName: "天ぷら"}
	Steak       = Category{searchWord: "steak", categoryName: "ステーキ"}
	Yoshoku     = Category{searchWord: "yoshoku", categoryName: "洋食"}
	Izakaya     = Category{searchWord: "izakaya", categoryName: "居酒屋"}
	Ramen       = Category{searchWord: "ramen_tokyo", categoryName: "ラーメン"}
	Yakiniku    = Category{searchWord: "yakiniku_tokyo", categoryName: "焼肉"}
	Cafe        = Category{searchWord: "cafe", categoryName: "カフェ"}
	Gyoza       = Category{searchWord: "gyoza", categoryName: "餃子"}
	Teishoku    = Category{searchWord: "teishoku", categoryName: "定食"}
	Kissaten    = Category{searchWord: "kissaten", categoryName: "喫茶店"}
	Pizza       = Category{searchWord: "pizza", categoryName: "ピザ"}
	Bistro      = Category{searchWord: "bistro", categoryName: "ビストロ"}
	Chinese     = Category{searchWord: "chinese_tokyo", categoryName: "中国料理"}
	Japanese    = Category{searchWord: "japanese_tokyo", categoryName: "日本料理"}
	French      = Category{searchWord: "french_tokyo", categoryName: "フレンチ"}
	Italian     = Category{searchWord: "italian_tokyo", categoryName: "イタリアン"}
	Bread       = Category{searchWord: "bread_tokyo", categoryName: "パン"}
	Sweets      = Category{searchWord: "sweets_tokyo", categoryName: "スイーツ"}
)

func CreateCategoryList() []Category {
	categoryList := make([]Category, 0)
	categoryList = append(categoryList, Curry)
	categoryList = append(categoryList, Unagi)
	categoryList = append(categoryList, Yakitori)
	categoryList = append(categoryList, Sushi)
	categoryList = append(categoryList, Tonkatsu)
	categoryList = append(categoryList, Soba)
	categoryList = append(categoryList, Hamburger)
	categoryList = append(categoryList, Okonomiyaki)
	categoryList = append(categoryList, Tempura)
	categoryList = append(categoryList, Steak)
	categoryList = append(categoryList, Yoshoku)
	categoryList = append(categoryList, Izakaya)
	categoryList = append(categoryList, Ramen)
	categoryList = append(categoryList, Yakiniku)
	categoryList = append(categoryList, Cafe)
	categoryList = append(categoryList, Gyoza)
	categoryList = append(categoryList, Teishoku)
	categoryList = append(categoryList, Kissaten)
	categoryList = append(categoryList, Pizza)
	categoryList = append(categoryList, Bistro)
	categoryList = append(categoryList, Chinese)
	categoryList = append(categoryList, Japanese)
	categoryList = append(categoryList, French)
	categoryList = append(categoryList, Italian)
	categoryList = append(categoryList, Bread)
	categoryList = append(categoryList, Sweets)

	return categoryList
}

func (c *Category) SearchWord() string {
	return c.searchWord
}

func (c *Category) CategoryName() string {
	return c.categoryName
}
