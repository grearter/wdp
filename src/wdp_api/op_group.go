package wdp_api

import "github.com/globalsign/mgo/bson"

type OpGroup struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
