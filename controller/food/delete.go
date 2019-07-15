package food

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"wdp/api"
	"wdp/dao"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, &api.FailResponse{ErrCode: ParaError, ErrMessage: "id无效"})
		return
	}

	err := dao.FoodDao.Delete(bson.ObjectIdHex(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
