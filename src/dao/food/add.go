package food

import (
	"github.com/globalsign/mgo/bson"
	"wdp_api"
)

func (food *Food) Add(foodApi *wdp_api.Food, id *bson.ObjectId) error {
	c := food.GetC(foodApi.Id)
	defer c.Database.Session.Close()

	if !foodApi.Id.Valid() {
		foodApi.Id = bson.NewObjectId()
	}

	*id = foodApi.Id
	return c.Insert(foodApi)
}
