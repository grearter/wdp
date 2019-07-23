package op_user

import (
	"github.com/globalsign/mgo/bson"
	"wdp_api"
)

func (u *User) Add(userApi *wdp_api.OpUser, id *bson.ObjectId) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(userApi)
	if err != nil {
		return
	}

	if id != nil {
		*id = userApi.Id
	}

	return
}
