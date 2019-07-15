package food

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"wdp/api"
	"wdp/dao"
	"wdp/dao/food"
)

func Update(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, &api.FailResponse{ErrCode: ParaError, ErrMessage: "id无效"})
		return
	}

	var f *food.Food

	if err := c.Bind(&f); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if f.Name == "" {
		c.JSON(http.StatusBadRequest, "name不能为空")
	}

	foodModel, err := dao.FoodDao.Update(bson.ObjectIdHex(id), f.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, foodModel)
	return
}
