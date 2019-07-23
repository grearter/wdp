package food

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"time"
	"wdp_api"
)

func Add(c *gin.Context) {
	var food *wdp_api.Food
	err := c.Bind(&food)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if food == nil || food.Name == "" {
		c.JSON(http.StatusBadRequest, &wdp_api.FailResponse{ErrCode: ParaError, ErrMessage: "name不能为空"})
		return
	}

	client, err := FoodDaoPool.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer client.Close()

	food.Id = bson.NewObjectId()
	food.Ut = time.Now().Unix()

	var id bson.ObjectId
	err = client.Call("Food.Add", food, &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
	return
}
