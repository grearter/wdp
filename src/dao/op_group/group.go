package op_group

import (
	"dao/mongo"
	"github.com/globalsign/mgo"
)

type Group struct {
}

func (*Group) Db() string {
	return "wdp"
}

func (*Group) Table() string {
	return "group"
}

func (g *Group) GetC() *mgo.Collection {
	db, table := g.Db(), g.Table()
	session := mongo.DBS[db]

	sessionCopy := session.Copy()
	c := sessionCopy.DB(db).C(table)
	return c
}

func (g *Group) GetTxnC() *mgo.Collection {
	db, table := g.Db(), "txn"
	session := mongo.DBS[db]

	sessionCopy := session.Copy()
	c := sessionCopy.DB(db).C(table)
	return c
}
