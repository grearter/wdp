package op_group

import (
	"dao/user"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

func (g *Group) LeaveGroup(groupId, userId bson.ObjectId) (err error) {
	tc := g.GetTxnC()
	defer tc.Database.Session.Close()

	ops := []txn.Op{{
		C:      g.Table(),
		Assert: txn.DocExists,
		Id:     groupId,
		Update: bson.M{
			"$pull": bson.M{
				"uids": userId,
			},
		},
	}, {
		C:      new(op_user.User).Table(),
		Assert: txn.DocExists,
		Id:     userId,
		Update: bson.M{
			"$pull": bson.M{
				"gids": groupId,
			},
		},
	}}

	runner := txn.NewRunner(tc)

	tid := bson.NewObjectId()
	err = runner.Run(ops, tid, nil)
	if err != nil {
		return
	}

	return
}
