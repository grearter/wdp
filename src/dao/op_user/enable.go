package op_user

import "github.com/globalsign/mgo/bson"

func (u *User) Enable(id bson.ObjectId, out interface{}) (err error) {
	c := u.GetC(id)
	defer c.Database.Session.Close()

	up := bson.M{
		"$set": bson.M{
			"disabled": false,
		},
	}

	return c.UpdateId(id, up)
}
