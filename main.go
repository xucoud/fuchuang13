package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func main() {
	InitData()
	c := gin.Default()

	funcMap := template.FuncMap{
		"pinc": func(i int) int {
			return res.PriceIndex[i] + 1
		},
		"vinc": func(i int) int {
			return res.VolumeIndex[i] + 1
		},
		"sinc": func(i int) int {
			return res.ShopIndex[i] + 1
		},
	}
	c.SetFuncMap(funcMap)
	c.Static("/js", "static/js/*")
	c.LoadHTMLGlob("view/*")

	c.GET("/show", getItemId)
	c.GET("/getPre", getPre)
	c.GET("/getNext", getNext)
	c.GET("/getVPre", getVPre)
	c.GET("/getVNext", getVNext)
	c.GET("/getSPre", getSPre)
	c.GET("/getSNext", getSNext)

	err := c.Run(":8088")
	if err != nil {
		log.Println("启动失败···")
		return
	}
}
