package food

import "github.com/globalsign/mgo/bson"

func (food *Food) Get(id bson.ObjectId, foodModel *Food) (err error) {
	c := food.GetC(id)
	defer c.Database.Session.Close()

	err = c.FindId(id).One(foodModel)
	return
}
