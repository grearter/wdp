package food

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (food *Food) Delete(id bson.ObjectId) (err error) {
	c := food.GetC()
	defer c.Database.Session.Close()

	err = c.RemoveId(id)
	if err == mgo.ErrNotFound {
		return
	}

	return
}
