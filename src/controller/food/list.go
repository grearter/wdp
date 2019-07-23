package food

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wdp_api"
)

func List(c *gin.Context) {

	client, err := FoodDaoPool.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer client.Close()

	var foods []*wdp_api.Food
	err = client.Call("Food.List", nil, &foods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, foods)
	return
}
