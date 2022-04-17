package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	res Page
	ps  = &PriceItem{}
	vs  = &VolumeItem{}
	si  = ShopItem{}
)

//首次加载
func getItemId(ctx *gin.Context) {
	ppage, err := strconv.Atoi(ctx.Query("ppage"))
	if err != nil || ppage <= 0 {
		ppage = 1
	}
	plimit, err := strconv.Atoi(ctx.Query("plimit"))
	if err != nil || plimit <= 0 {
		plimit = 10
	}
	vpage, err := strconv.Atoi(ctx.Query("vpage"))
	if err != nil || vpage <= 0 {
		vpage = 1
	}
	vlimit, err := strconv.Atoi(ctx.Query("vlimit"))
	if err != nil || vlimit <= 0 {
		vlimit = 10
	}
	spage, err := strconv.Atoi(ctx.Query("spage"))
	if err != nil || spage <= 0 {
		spage = 1
	}
	slimit, err := strconv.Atoi(ctx.Query("slimit"))
	if err != nil || slimit <= 0 {
		slimit = 10
	}
	pg := len(Price)/plimit + 1
	if pg < ppage {
		ppage = pg
	}
	vg := len(Volume)/vlimit + 1
	if vg < vpage {
		vpage = vg
	}
	sg := len(SD)/slimit + 1
	if sg < spage {
		spage = sg
	}
	pd, pi := ps.getData(ppage, plimit)
	vd, vi := vs.getData(vpage, vlimit)
	sd, sid := si.getData(spage, slimit)

	pc, _ := json.Marshal(PriceData)
	vc, _ := json.Marshal(VolumeData)
	ac, _ := json.Marshal(All)
	res = Page{
		PricePageNum:      ppage,
		PriceLimit:        plimit,
		PriceData:         pd,
		PriceIndex:        pi,
		PricePage:         pg,
		PriceCateNameLv1:  pc,
		VolumePageNum:     vpage,
		VolumeLimit:       vlimit,
		VolumeData:        vd,
		VolumeIndex:       vi,
		VolumePage:        vg,
		VolumeCateNameLv1: vc,
		AllCateNameLv1:    ac,
		ShopData:          sd,
		ShopIndex:         sid,
		ShopPageNum:       spage,
		ShopPage:          sg,
		ShopLimit:         slimit,
	}
	ctx.HTML(200, "show.html", res)
}

//上一页
func getPre(ctx *gin.Context) {
	ppage, err := strconv.Atoi(ctx.Query("ppage"))
	if err != nil || ppage <= 1 {
		ppage = 1
	} else {
		ppage--
	}
	plimit, err := strconv.Atoi(ctx.Query("plimit"))
	if err != nil || plimit <= 0 {
		plimit = 10
	}
	vpage, err := strconv.Atoi(ctx.Query("vpage"))
	if err != nil || vpage <= 1 {
		vpage = 1
	}
	vlimit, err := strconv.Atoi(ctx.Query("vlimit"))
	if err != nil || vlimit <= 0 {
		vlimit = 10
	}

	spage, err := strconv.Atoi(ctx.Query("spage"))
	if err != nil || spage <= 0 {
		spage = 1
	}
	slimit, err := strconv.Atoi(ctx.Query("slimit"))
	if err != nil || slimit <= 0 {
		slimit = 10
	}
	pg := len(Price)/plimit + 1
	if pg < ppage {
		ppage = pg
	}
	vg := len(Volume)/vlimit + 1
	if vg < vpage {
		vpage = vg
	}
	sg := len(SD)/slimit + 1
	if sg < spage {
		spage = sg
	}
	pd, pi := ps.getData(ppage, plimit)
	vd, vi := vs.getData(vpage, vlimit)
	sd, sid := si.getData(spage, slimit)

	pc, _ := json.Marshal(PriceData)
	vc, _ := json.Marshal(VolumeData)
	ac, _ := json.Marshal(All)
	res = Page{
		PricePageNum:      ppage,
		PriceLimit:        plimit,
		PriceData:         pd,
		PriceIndex:        pi,
		PricePage:         pg,
		PriceCateNameLv1:  pc,
		VolumePageNum:     vpage,
		VolumeLimit:       vlimit,
		VolumeData:        vd,
		VolumeIndex:       vi,
		VolumePage:        vg,
		VolumeCateNameLv1: vc,
		AllCateNameLv1:    ac,
		ShopData:          sd,
		ShopIndex:         sid,
		ShopPageNum:       spage,
		ShopPage:          sg,
		ShopLimit:         slimit,
	}
	ctx.HTML(200, "show.html", res)
}

//下一页
func getNext(ctx *gin.Context) {
	ppage, err := strconv.Atoi(ctx.Query("ppage"))
	if err != nil || ppage <= 0 {
		ppage = 1
	} else {
		ppage++
	}
	plimit, err := strconv.Atoi(ctx.Query("plimit"))
	if err != nil || plimit <= 0 {
		plimit = 10
	}
	vpage, err := strconv.Atoi(ctx.Query("vpage"))
	if err != nil || vpage <= 0 {
		vpage = 1
	}
	vlimit, err := strconv.Atoi(ctx.Query("vlimit"))
	if err != nil || vlimit <= 0 {
		vlimit = 10
	}

	spage, err := strconv.Atoi(ctx.Query("spage"))
	if err != nil || spage <= 0 {
		spage = 1
	}
	slimit, err := strconv.Atoi(ctx.Query("slimit"))
	if err != nil || slimit <= 0 {
		slimit = 10
	}
	pg := len(Price)/plimit + 1
	if pg < ppage {
		ppage = pg
	}
	vg := len(Volume)/vlimit + 1
	if vg < vpage {
		vpage = vg
	}
	sg := len(SD)/slimit + 1
	if sg < spage {
		spage = sg
	}
	pd, pi := ps.getData(ppage, plimit)
	vd, vi := vs.getData(vpage, vlimit)
	sd, sid := si.getData(spage, slimit)

	pc, _ := json.Marshal(PriceData)
	vc, _ := json.Marshal(VolumeData)
	ac, _ := json.Marshal(All)
	res = Page{
		PricePageNum:      ppage,
		PriceLimit:        plimit,
		PriceData:         pd,
		PriceIndex:        pi,
		PricePage:         pg,
		PriceCateNameLv1:  pc,
		VolumePageNum:     vpage,
		VolumeLimit:       vlimit,
		VolumeData:        vd,
		VolumeIndex:       vi,
		VolumePage:        vg,
		VolumeCateNameLv1: vc,
		AllCateNameLv1:    ac,
		ShopData:          sd,
		ShopIndex:         sid,
		ShopPageNum:       spage,
		ShopPage:          sg,
		ShopLimit:         slimit,
	}
	ctx.HTML(200, "show.html", res)
}

//上一页
func getVPre(ctx *gin.Context) {
	ppage, err := strconv.Atoi(ctx.Query("ppage"))
	if err != nil || ppage <= 0 {
		ppage = 1
	}
	plimit, err := strconv.Atoi(ctx.Query("plimit"))
	if err != nil || plimit <= 0 {
		plimit = 10
	}
	vpage, err := strconv.Atoi(ctx.Query("vpage"))
	if err != nil || vpage <= 1 {
		vpage = 1
	} else {
		vpage--
	}
	vlimit, err := strconv.Atoi(ctx.Query("vlimit"))
	if err != nil || vlimit <= 0 {
		vlimit = 10
	}

	spage, err := strconv.Atoi(ctx.Query("spage"))
	if err != nil || spage <= 0 {
		spage = 1
	}
	slimit, err := strconv.Atoi(ctx.Query("slimit"))
	if err != nil || slimit <= 0 {
		slimit = 10
	}
	pg := len(Price)/plimit + 1
	if pg < ppage {
		ppage = pg
	}
	vg := len(Volume)/vlimit + 1
	if vg < vpage {
		vpage = vg
	}
	sg := len(SD)/slimit + 1
	if sg < spage {
		spage = sg
	}
	pd, pi := ps.getData(ppage, plimit)
	vd, vi := vs.getData(vpage, vlimit)
	sd, sid := si.getData(spage, slimit)

	pc, _ := json.Marshal(PriceData)
	vc, _ := json.Marshal(VolumeData)
	ac, _ := json.Marshal(All)
	res = Page{
		PricePageNum:      ppage,
		PriceLimit:        plimit,
		PriceData:         pd,
		PriceIndex:        pi,
		PricePage:         pg,
		PriceCateNameLv1:  pc,
		VolumePageNum:     vpage,
		VolumeLimit:       vlimit,
		VolumeData:        vd,
		VolumeIndex:       vi,
		VolumePage:        vg,
		VolumeCateNameLv1: vc,
		AllCateNameLv1:    ac,
		ShopData:          sd,
		ShopIndex:         sid,
		ShopPageNum:       spage,
		ShopPage:          sg,
		ShopLimit:         slimit,
	}
	ctx.HTML(200, "show.html", res)
}

//下一页
func getVNext(ctx *gin.Context) {
	ppage, err := strconv.Atoi(ctx.Query("ppage"))
	if err != nil || ppage <= 0 {
		ppage = 1
	}
	plimit, err := strconv.Atoi(ctx.Query("plimit"))
	if err != nil || plimit <= 0 {
		plimit = 10
	}
	vpage, err := strconv.Atoi(ctx.Query("vpage"))
	if err != nil || vpage <= 0 {
		vpage = 1
	} else {
		vpage++
	}
	vlimit, err := strconv.Atoi(ctx.Query("vlimit"))
	if err != nil || vlimit <= 0 {
		vlimit = 10
	}

	spage, err := strconv.Atoi(ctx.Query("spage"))
	if err != nil || spage <= 0 {
		spage = 1
	}
	slimit, err := strconv.Atoi(ctx.Query("slimit"))
	if err != nil || slimit <= 0 {
		slimit = 10
	}
	pg := len(Price)/plimit + 1
	if pg < ppage {
		ppage = pg
	}
	vg := len(Volume)/vlimit + 1
	if vg < vpage {
		vpage = vg
	}
	sg := len(SD)/slimit + 1
	if sg < spage {
		spage = sg
	}
	pd, pi := ps.getData(ppage, plimit)
	vd, vi := vs.getData(vpage, vlimit)
	sd, sid := si.getData(spage, slimit)

	pc, _ := json.Marshal(PriceData)
	vc, _ := json.Marshal(VolumeData)
	ac, _ := json.Marshal(All)
	res = Page{
		PricePageNum:      ppage,
		PriceLimit:        plimit,
		PriceData:         pd,
		PriceIndex:        pi,
		PricePage:         pg,
		PriceCateNameLv1:  pc,
		VolumePageNum:     vpage,
		VolumeLimit:       vlimit,
		VolumeData:        vd,
		VolumeIndex:       vi,
		VolumePage:        vg,
		VolumeCateNameLv1: vc,
		AllCateNameLv1:    ac,
		ShopData:          sd,
		ShopIndex:         sid,
		ShopPageNum:       spage,
		ShopPage:          sg,
		ShopLimit:         slimit,
	}
	ctx.HTML(200, "show.html", res)
}

//上一页
func getSPre(ctx *gin.Context) {
	ppage, err := strconv.Atoi(ctx.Query("ppage"))
	if err != nil || ppage <= 1 {
		ppage = 1
	}
	plimit, err := strconv.Atoi(ctx.Query("plimit"))
	if err != nil || plimit <= 0 {
		plimit = 10
	}
	vpage, err := strconv.Atoi(ctx.Query("vpage"))
	if err != nil || vpage <= 1 {
		vpage = 1
	}
	vlimit, err := strconv.Atoi(ctx.Query("vlimit"))
	if err != nil || vlimit <= 0 {
		vlimit = 10
	}

	spage, err := strconv.Atoi(ctx.Query("spage"))
	if err != nil || spage <= 1 {
		spage = 1
	} else {
		spage--
	}
	slimit, err := strconv.Atoi(ctx.Query("slimit"))
	if err != nil || slimit <= 0 {
		slimit = 10
	}
	pg := len(Price)/plimit + 1
	if pg < ppage {
		ppage = pg
	}
	vg := len(Volume)/vlimit + 1
	if vg < vpage {
		vpage = vg
	}
	sg := len(SD)/slimit + 1
	if sg < spage {
		spage = sg
	}
	pd, pi := ps.getData(ppage, plimit)
	vd, vi := vs.getData(vpage, vlimit)
	sd, sid := si.getData(spage, slimit)

	pc, _ := json.Marshal(PriceData)
	vc, _ := json.Marshal(VolumeData)
	ac, _ := json.Marshal(All)
	res = Page{
		PricePageNum:      ppage,
		PriceLimit:        plimit,
		PriceData:         pd,
		PriceIndex:        pi,
		PricePage:         pg,
		PriceCateNameLv1:  pc,
		VolumePageNum:     vpage,
		VolumeLimit:       vlimit,
		VolumeData:        vd,
		VolumeIndex:       vi,
		VolumePage:        vg,
		VolumeCateNameLv1: vc,
		AllCateNameLv1:    ac,
		ShopData:          sd,
		ShopIndex:         sid,
		ShopPageNum:       spage,
		ShopPage:          sg,
		ShopLimit:         slimit,
	}
	ctx.HTML(200, "show.html", res)
}

//下一页
func getSNext(ctx *gin.Context) {
	ppage, err := strconv.Atoi(ctx.Query("ppage"))
	if err != nil || ppage <= 0 {
		ppage = 1
	}
	plimit, err := strconv.Atoi(ctx.Query("plimit"))
	if err != nil || plimit <= 0 {
		plimit = 10
	}
	vpage, err := strconv.Atoi(ctx.Query("vpage"))
	if err != nil || vpage <= 0 {
		vpage = 1
	}
	vlimit, err := strconv.Atoi(ctx.Query("vlimit"))
	if err != nil || vlimit <= 0 {
		vlimit = 10
	}

	spage, err := strconv.Atoi(ctx.Query("spage"))
	if err != nil || spage <= 0 {
		spage = 1
	} else {
		spage++
	}
	slimit, err := strconv.Atoi(ctx.Query("slimit"))
	if err != nil || slimit <= 0 {
		slimit = 10
	}
	pg := len(Price)/plimit + 1
	if pg < ppage {
		ppage = pg
	}
	vg := len(Volume)/vlimit + 1
	if vg < vpage {
		vpage = vg
	}
	sg := len(SD)/slimit + 1
	if sg < spage {
		spage = sg
	}
	pd, pi := ps.getData(ppage, plimit)
	vd, vi := vs.getData(vpage, vlimit)
	sd, sid := si.getData(spage, slimit)

	pc, _ := json.Marshal(PriceData)
	vc, _ := json.Marshal(VolumeData)
	ac, _ := json.Marshal(All)
	res = Page{
		PricePageNum:      ppage,
		PriceLimit:        plimit,
		PriceData:         pd,
		PriceIndex:        pi,
		PricePage:         pg,
		PriceCateNameLv1:  pc,
		VolumePageNum:     vpage,
		VolumeLimit:       vlimit,
		VolumeData:        vd,
		VolumeIndex:       vi,
		VolumePage:        vg,
		VolumeCateNameLv1: vc,
		AllCateNameLv1:    ac,
		ShopData:          sd,
		ShopIndex:         sid,
		ShopPageNum:       spage,
		ShopPage:          sg,
		ShopLimit:         slimit,
	}
	ctx.HTML(200, "show.html", res)
}
