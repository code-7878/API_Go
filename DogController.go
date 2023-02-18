package main

import "github.com/gin-gonic/gin"

func GetAllDogs(c *gin.Context) {
	c.JSON(200, FindAllDogs())
}

//довавить котика

func AddDog(c *gin.Context) {
	var dog Dog
	if err := c.BindJSON(&dog); err != nil {
		return
	}
	c.JSON(200, CreateDog(dog))
}

func GetDog(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, FindDogById(id))
}

func DeleteDog(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, DropDogById(id))
}

func EditDog(c *gin.Context) {
	id := c.Param("id")

	var dog Dog
	if err := c.BindJSON(&dog); err != nil {
		return
	}

	dog.ID = id

	c.JSON(200, UpdateCat(dog))
}
