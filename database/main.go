package main

import (
	"context"
	"encoding/json"
	"flag"
	"geo/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"time"
)

var (
	scriptFilePath = flag.String("script", "database/json/region.json", "脚本数据路径")
	regionList     []*domain.GeoNode
)

func main() {
	flag.Parse()
	if len(*scriptFilePath) == 0 {
		panic("脚本不存在!")
	}
	//解析json文件
	data, err := ioutil.ReadFile(*scriptFilePath)
	if err != nil {
		log.Printf("读取文件失败:%v", err)
		return
	}
	var script = &domain.ChinaLocation{}
	if err := json.Unmarshal(data, script); err != nil {
		log.Printf("解析json失败:%v", err)
		return
	}
	//整理入库数据结构
	//仅保存每个地区信息，地区信息有对应的城市和省份
	MakeRegionList(script.Node)
	//入库
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opt := options.Client()
	//自己的mongo配置
	mongoDSN := "mongodb://xxx"
	client, err := mongo.Connect(ctx,
		opt.ApplyURI(mongoDSN),
	)
	if err != nil {
		log.Println(err)
		return
	}
	collection := client.Database("xxx").Collection("xxx")

	for _, v := range regionList {
		_, err = collection.InsertOne(context.Background(), v)
		if err != nil {
			log.Printf("插入mongo失败:%v %v", v, err)
		}
	}
	log.Printf("处理完毕")
}

func Node2GeoNode(node *domain.Node) *domain.GeoNode {
	return &domain.GeoNode{
		Name: node.Name,
		Center: []float64{
			node.Center.Longitude,
			node.Center.Latitude,
		},
		Country:  node.Country,
		Province: node.Province,
		City:     node.City,
		Level:    node.Level,
	}
}
 //递归处理每个地区信息，如果是level=省或市，把值设置到下一个level中
func MakeRegionList(node *domain.Node) {
	if node.Level == "district" {
		regionList = append(regionList, Node2GeoNode(node))
	} else {
		if len(node.Districts) > 0 {
			for _, v := range node.Districts {
				if node.Level == "country" {
					v.Country = node.Name
				} else if node.Level == "province" {
					v.Country = node.Country
					v.Province = node.Name

				} else if node.Level == "city" {
					v.City = node.Name
					v.Country = node.Country
					v.Province = node.Province
				}
				MakeRegionList(v)
			}
		}
	}
}
