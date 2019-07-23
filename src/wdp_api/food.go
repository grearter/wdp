package wdp_api

import "github.com/globalsign/mgo/bson"

type Food struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Ut   int64         `json:"ut" bson:"ut"`
}
