package food

import "github.com/globalsign/mgo"

func (food *Food) List() (foodModels []*Food, err error) {
	c := food.GetC()
	defer c.Database.Session.Close()

	err = c.Find(nil).Sort("-ut").All(&foodModels)
	if err == mgo.ErrNotFound {
		err = nil
	}

	return
}
