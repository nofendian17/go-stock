package entity

type Broker struct {
	Code    string `bson:"code"`
	Name    string `bson:"name"`
	License string `bson:"license"`
}
