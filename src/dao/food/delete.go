package food

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"wdp_api"
)

func (food *Food) Delete(id bson.ObjectId, output *wdp_api.Food) (err error) {
	c := food.GetC(id)
	defer c.Database.Session.Close()

	ch := mgo.Change{
		Remove: true,
	}

	_, err = c.FindId(id).Apply(ch, output)
	return
}
