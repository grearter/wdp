package food

import (
	"github.com/globalsign/mgo"
	"wdp_api"
)

func (food *Food) List(input interface{}, foodModels *[]*wdp_api.Food) (err error) {
	c := food.GetC(nil)
	defer c.Database.Session.Close()

	err = c.Find(nil).Sort("-ut").All(foodModels)

	if err == mgo.ErrNotFound {
		err = nil
	}

	return
}
