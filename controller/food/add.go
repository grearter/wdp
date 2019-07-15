package food

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"time"
	"wdp/api"
	"wdp/dao"
	"wdp/dao/food"
)

func Add(c *gin.Context) {
	var f food.Food
	err := c.Bind(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("food: %v\b", f)
	if f.Name == "" {
		c.JSON(http.StatusBadRequest, &api.FailResponse{ErrCode: ParaError, ErrMessage: "name不能为空"})
		return
	}

	foodModel := &food.Food{
		Id:   bson.NewObjectId(),
		Name: f.Name,
		Ut:   time.Now().Unix(),
	}

	err = dao.FoodDao.Add(foodModel)
	if err != nil {
		if mgo.IsDup(err) {
			c.JSON(http.StatusBadRequest, &api.FailResponse{ErrCode: DupError, ErrMessage: "name不能重复"})
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, foodModel)
	return
}
