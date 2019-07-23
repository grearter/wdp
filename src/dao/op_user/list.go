package op_user

import "wdp_api"

func (u *User) List(in interface{}, userApis *[]*wdp_api.OpUser) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()

	return c.Find(nil).All(userApis)
}
