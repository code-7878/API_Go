package main

import "github.com/gin-gonic/gin"

// GetAllCats получить из списка всех котиков.
func GetAllCats(c *gin.Context) {
	c.JSON(200, FindAllCats())
}

//довавить котика

func AddCat(c *gin.Context) {
	var cat Cat
	if err := c.BindJSON(&cat); err != nil {
		return
	}
	c.JSON(200, CreateCat(cat))
}

func GetCat(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, FindCatById(id))
}

func DeleteCat(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, DropCatById(id))
}

func EditCat(c *gin.Context) {
	id := c.Param("id")

	var cat Cat
	if err := c.BindJSON(&cat); err != nil {
		return
	}
	
	cat.ID = id

	c.JSON(200, UpdateCat(cat))
}