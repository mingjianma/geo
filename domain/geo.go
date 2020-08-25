package domain

type ChinaLocation struct {
	*Node
}

type Node struct {
	Name      string  `json:"name"`
	Center    *Center `json:"center"`
	Country   string  `json:"country"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	Level     string  `json:"level"`
	Districts []*Node `json:"districts"`
}

type Center struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type GeoNode struct {
	Name     string    `bson:"name"`
	Center   []float64 `bson:"center"`
	Country  string    `bson:"country"`
	Province string    `bson:"province"`
	City     string    `bson:"city"`
	Level    string    `bson:"level"`
}

type GeoMongoResult struct {
	Ok      int          `bson:"ok"`
	Results []*MapResult `bson:"results"`
	Stats   *MongoStatus `bson:"stats"`
}

type MapResult struct {
	Dis float64 `bson:"dis"`
	Obj *GeoObj `bson:"obj"`
}

type GeoObj struct {
	ID       string    `bson:"_id"`
	Name     string    `bson:"name"`
	Center   []float64 `bson:"center"`
	Country  string    `bson:"country"`
	Province string    `bson:"province"`
	City     string    `bson:"city"`
	Level    string    `bson:"level"`
}

type MongoStatus struct {
	AvgDistance   float64 `bson:"avgDistance"`
	MaxDistance   float64 `bson:"max_distance"`
	Nscanned      int     `bson:"nscanned"`
	ObjectsLoaded int     `bson:"objectsLoaded"`
	Time          int     `bson:"time"`
}
