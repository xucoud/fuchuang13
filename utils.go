package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var enc mahonia.Decoder

var (
	Price        []string
	Volume       []string
	AllData      = make(map[string]*Commodity)
	PriceData    = make(map[string]int)
	VolumeData   = make(map[string]int)
	All          = make(map[string]int)
	ShopData     = make(map[string]int)
	shopNameToId = make(map[string]string)
	SD           = make([]Shop, 0)
)

const (
	data06 = "./data/data_202106.tsv"
	data07 = "./data/data_202107.tsv"
	data08 = "./data/data_202108.tsv"
	data09 = "./data/data_202109.tsv"
	final  = "./data/result.txt"
)

// InitData 初始化数据
func InitData() {
	InitAllCom()
	InitAbnormal()
	getChar()
}

// InitAbnormal 初始化销量异常商品数据
func InitAbnormal() {
	f, err := os.Open(final)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer f.Close()
	err = readTxt(f)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}

// InitAllCom 初始化所有商品数据
func InitAllCom() {
	readTsv(data06)
	readTsv(data07)
	readTsv(data08)
	readTsv(data09)
}

//按行读取TXT文件
func readTxt(r io.Reader) error {
	reader := bufio.NewReader(r)

	//l := make([]string, 0, 64)
	// 按行读取
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		tmp := strings.Split(string(line), ",")
		if tmp[1] == "价格异常" {
			Price = append(Price, tmp[0])
		} else if tmp[1] == "销量异常" {
			Volume = append(Volume, tmp[0])
		}
	}

	return nil
}

//读取tsv文件
func readTsv(fileName string) {
	enc = mahonia.NewDecoder("gbk")
	start1 := time.Now()
	FileHandle, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()

	lineReader := bufio.NewReader(FileHandle)
	for {
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		tmp := strings.Split(string(line), "\t")
		AllData[tmp[1]] = &Commodity{
			ItemId:      tmp[1],
			ItemName:    enc.ConvertString(tmp[2]),
			CateNameLv1: enc.ConvertString(tmp[8]),
			CateNameLv2: enc.ConvertString(tmp[9]),
			CateNameLv3: enc.ConvertString(tmp[10]),
			UserId:      enc.ConvertString(tmp[19]),
			ShopName:    enc.ConvertString(tmp[20]),
		}
		shopNameToId[AllData[tmp[1]].UserId] = AllData[tmp[1]].ShopName
	}
	fmt.Println("readEachLineReader spend : ", time.Now().Sub(start1))
}

//处理图标数据
func getChar() {
	//color := []string{"0780cf", "765005", "fa6d1d", "0e2c82", "b6b51f", "da1f18",
	//	"701866", "f47a75", "009db2", "024b51", "0780cf", "765005", "3682be", "45a776",
	//	"f05326", "eed777", "334f65", "b3974e", "38cb7d", "ddae33", "844bb3", "93c555", "5f6694", "df3881"}

	for _, v := range Price {
		AllData[v].Type = "价格异常"
		if _, ok := PriceData[AllData[v].CateNameLv1]; ok {
			PriceData[AllData[v].CateNameLv1]++
		} else {
			PriceData[AllData[v].CateNameLv1] = 1
		}
		if _, ok := All[AllData[v].CateNameLv1]; ok {
			All[AllData[v].CateNameLv1]++
		} else {
			All[AllData[v].CateNameLv1] = 1
		}

		if _, ok := ShopData[AllData[v].UserId]; ok {
			ShopData[AllData[v].UserId]++
		} else {
			ShopData[AllData[v].UserId] = 1
		}
	}
	for _, v := range Volume {
		AllData[v].Type = "销量异常"
		if _, ok := VolumeData[AllData[v].CateNameLv1]; ok {
			VolumeData[AllData[v].CateNameLv1]++
		} else {
			VolumeData[AllData[v].CateNameLv1] = 1
		}
		if _, ok := All[AllData[v].CateNameLv1]; ok {
			All[AllData[v].CateNameLv1]++
		} else {
			All[AllData[v].CateNameLv1] = 1
		}

		if _, ok := ShopData[AllData[v].UserId]; ok {
			ShopData[AllData[v].UserId]++
		} else {
			ShopData[AllData[v].UserId] = 1
		}
	}

	for k, v1 := range ShopData {
		if k == "\\n" || k == "" {
			continue
		}
		t := Shop{
			Id:   k,
			Name: shopNameToId[k],
			Num:  v1,
		}
		SD = append(SD, t)
		//if v, ok := AllData[k]; ok {
		//	t := Shop{
		//		Id:   v.UserId,
		//		Name: v.ShopName,
		//		Num:  v1,
		//	}
		//	SD = append(SD, t)
		//}
	}
	sort.Slice(SD, func(i, j int) bool {
		return SD[i].Num > SD[j].Num
	})

	//fmt.Println(VolumeData)
	//fmt.Println(PriceData)
	//fmt.Println(All)
}
