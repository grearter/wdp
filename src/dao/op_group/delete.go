package op_group

import (
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

func (g *Group) Delete(id bson.ObjectId, out interface{}) (err error) {
	tc := g.GetTxnC()
	defer tc.Database.Session.Close()

	ops := []txn.Op{{
		C: g.Table(),
		Assert: bson.M{
			"uids": bson.M{
				"$size": 0,
			},
		},
		Id:     id,
		Remove: true,
	}}

	tid := bson.NewObjectId()
	runner := txn.NewRunner(tc)

	err = runner.Run(ops, tid, nil)
	if err != nil {
		return
	}

	return
}
