package food

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"wdp/mongo"
)

type Food struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Ut   int64         `json:"ut" bson:"ut"`
}

func (*Food) Db() string {
	return "wdp"
}

func (*Food) Table() string {
	return "food"
}

func (food *Food) GetC() *mgo.Collection {
	db, table := food.Db(), food.Table()
	session := mongo.DBS[db]

	sessionCopy := session.Copy()
	c := sessionCopy.DB(db).C(table)
	return c
}
