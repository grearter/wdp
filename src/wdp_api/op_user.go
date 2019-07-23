package wdp_api

import "github.com/globalsign/mgo/bson"

type OpUser struct {
	Id       bson.ObjectId   `json:"id" bson:"_id"`
	Name     string          `json:"name" bson:"name"`
	GroupIds []bson.ObjectId `json:"group_ids" bson:"gids"`
	Disabled bool            `json:"disabled" bson:"disabled"`
	Secret   string          `json:"secret" bson:"secret"`
}
