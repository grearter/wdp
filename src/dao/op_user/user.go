package op_user

import (
	"dao/mongo"
	"github.com/globalsign/mgo"
)

type User struct {
}

func (*User) Db() string {
	return "wdp"
}

func (*User) Table() string {
	return "user"
}

func (u *User) GetC() *mgo.Collection {
	db, table := u.Db(), u.Table()
	session := mongo.DBS[db]

	sessionCopy := session.Copy()
	c := sessionCopy.DB(db).C(table)
	return c
}

func (u *User) GetTxnC() *mgo.Collection {
	db, table := u.Db(), "txn"
	session := mongo.DBS[db]

	sessionCopy := session.Copy()
	c := sessionCopy.DB(db).C(table)
	return c
}
