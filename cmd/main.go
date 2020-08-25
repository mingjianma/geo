package main

import (
	"context"
	"encoding/json"
	"geo/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	longitude, latitude := 113.290717, 22.767033
	filter := bson.D{
		{"geoNear", "region"},
		{"spherical", true},
		{"near", [2]float64{
			longitude, latitude},
		},
		{"num", 1},
	}
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
	db := client.Database("xxx")
	//db.runCommand({geoNear:'region', near:[113.290717,22.767033], spherical:true});
	res := db.RunCommand(ctx, filter)
	location := &domain.GeoMongoResult{}
	err = res.Decode(location)
	str, _ := json.Marshal(location)
	//广东省佛山市顺德区
	log.Printf("%v %s", err, string(str))
}
