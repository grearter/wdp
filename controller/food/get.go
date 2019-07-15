package food

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"wdp/api"
	"wdp/dao"
)

func Get(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, &api.FailResponse{ErrCode: ParaError, ErrMessage: "id无效"})
		return
	}

	foodModel, err := dao.FoodDao.Get(bson.ObjectIdHex(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, foodModel)
	return
}
