package main

type PriceItem struct{}

func (p *PriceItem) getData(page, limit int) ([]Commodity, []int) {
	start := page*limit - limit
	end := page * limit
	if end > len(Price) {
		end = len(Price)
	}
	tmp := make([]int, 0)
	for i := start; i < end; i++ {
		tmp = append(tmp, i)
	}
	ids := Price[start:end]
	data := make([]Commodity, len(ids))
	for i, v := range ids {
		data[i] = *AllData[v]
	}

	return data, tmp
}

type VolumeItem struct{}

func (p *VolumeItem) getData(page, limit int) ([]Commodity, []int) {
	start := page*limit - limit
	end := page * limit
	if end > len(Volume) {
		end = len(Volume)
	}
	tmp := make([]int, 0)
	for i := start; i < end; i++ {
		tmp = append(tmp, i)
	}
	ids := Volume[start:end]
	data := make([]Commodity, len(ids))
	for i, v := range ids {
		data[i] = *AllData[v]
	}

	return data, tmp
}

type ShopItem struct{}

func (s ShopItem) getData(page, limit int) ([]Shop, []int) {
	start := page*limit - limit
	end := page * limit
	if end > len(ShopData) {
		end = len(ShopData)
	}
	tmp := make([]int, 0)
	for i := start; i < end; i++ {
		tmp = append(tmp, i)
	}
	data := SD[start:end]

	return data, tmp
}
