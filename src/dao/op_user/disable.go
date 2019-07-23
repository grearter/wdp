package op_user

import "github.com/globalsign/mgo/bson"

func (u *User) Disable(id bson.ObjectId, out interface{}) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()

	up := bson.M{
		"$set": bson.M{
			"disabled": true,
		},
	}

	return c.UpdateId(id, up)
}
