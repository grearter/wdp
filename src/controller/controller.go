package main

import (
	"controller/food"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	route(r)

	err := r.Run()
	if err != nil {
		fmt.Printf("run err: %v\n", err)
	}

	return
}

func route(r *gin.Engine) {
	r.POST("/food/add", food.Add)
	//r.GET("/food/get/:id", food.Get)
	//r.PUT("/food/update/:id", food.Update)
	//r.DELETE("food/delete/:id", food.Delete)
	r.GET("/food/list", food.List)
}
