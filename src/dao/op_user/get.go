package op_user

import (
	"github.com/globalsign/mgo/bson"
	"wdp_api"
)

func (u *User) Get(id bson.ObjectId, userApi *wdp_api.OpUser) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()

	return c.FindId(id).One(userApi)
}
