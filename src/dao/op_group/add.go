package op_group

import (
	"github.com/globalsign/mgo/bson"
	"wdp_api"
)

func (g *Group) Add(groupApi *wdp_api.OpGroup, id *bson.ObjectId) (err error) {
	c := g.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(groupApi)
	if err != nil {
		return
	}

	if id != nil {
		*id = groupApi.Id
	}

	return
}
