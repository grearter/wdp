package op_group

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"wdp_api"
)

func (g *Group) Get(id bson.ObjectId, groupApi *wdp_api.OpGroup) error {
	c := g.GetC()
	defer c.Database.Session.Close()

	return c.FindId(id).One(groupApi)
}

func (g *Group) Gets(ids []bson.ObjectId, groupApis *[]*wdp_api.OpGroup) (err error) {
	c := g.GetC()
	defer c.Database.Session.Close()

	q := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}

	err = c.Find(q).All(groupApis)

	if err == mgo.ErrNotFound {
		err = nil
		*groupApis = make([]*wdp_api.OpGroup, 0)
	}

	return
}
