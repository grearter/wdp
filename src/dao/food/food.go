package food

import (
	"dao/mongo"
	"github.com/globalsign/mgo"
)

type Food struct {
}

func (*Food) Db(id interface{}) string {
	return "wdp"
}

func (*Food) Table(id interface{}) string {
	return "food"
}

func (food *Food) GetC(id interface{}) *mgo.Collection {
	db, table := food.Db(id), food.Table(id)
	session := mongo.DBS[db]

	sessionCopy := session.Copy()
	c := sessionCopy.DB(db).C(table)
	return c
}
