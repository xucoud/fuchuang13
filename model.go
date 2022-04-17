package main

type Commodity struct {
	ItemId      string
	ItemName    string
	CateNameLv1 string
	CateNameLv2 string
	CateNameLv3 string
	ShopName    string
	Type        string
}

type Shop struct {
	Name string
	Num  int
}

type Page struct {
	PricePageNum      int
	PriceLimit        int
	PriceData         []Commodity
	PriceIndex        []int
	PricePage         int
	PriceCateNameLv1  []byte
	VolumePageNum     int
	VolumeLimit       int
	VolumeData        []Commodity
	VolumeIndex       []int
	VolumePage        int
	VolumeCateNameLv1 []byte
	AllCateNameLv1    []byte
	ShopPageNum       int
	ShopLimit         int
	ShopData          []Shop
	ShopIndex         []int
	ShopPage          int
}
