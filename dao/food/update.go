package food

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

func (food *Food) Update(id bson.ObjectId, name string) (foodModel *Food, err error) {
	c := food.GetC()
	defer c.Database.Session.Close()

	ch := mgo.Change{

		Update: bson.M{
			"name": name,
			"ut":   time.Now().Unix(),
		},
		ReturnNew: true,
	}

	_, err = c.FindId(id).Apply(ch, &foodModel)
	return
}
