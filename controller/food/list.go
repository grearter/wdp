package food

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wdp/api"
	"wdp/dao"
	"wdp/dao/food"
)

func List(c *gin.Context) {
	foods, err := dao.FoodDao.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &api.FailResponse{ErrCode: SrvError, ErrMessage: err.Error()})
		return
	}

	if foods == nil {
		foods = make([]*food.Food, 0)
	}

	c.JSON(http.StatusOK, foods)
	return
}
