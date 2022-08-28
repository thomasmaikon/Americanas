package controller

import (
	"projeto/Americanas/model"
	"projeto/Americanas/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var servicePlanet = service.ServicePlanet{}

func CreatePlanet(c *gin.Context) {
	var planet model.PLanet
	c.BindJSON(&planet)

	if servicePlanet.CreateNewPlanet(planet) {
		c.JSON(201, gin.H{
			"Info": "planeta cadastrado com sucesso",
		})
		return
	} else {
		c.JSON(409, gin.H{
			"Info": "Planeta ja existente",
		})
	}
}

func ListPlanets(c *gin.Context) {
	findByName := c.Query("name")
	findById := c.Query("id")

	id, _ := strconv.ParseInt(findById, 10, 64)

	c.JSON(200, gin.H{
		"Info": servicePlanet.Find(findByName, id),
	})
}

func RemovePlanet(c *gin.Context) {
	name := c.Query("name")

	c.JSON(200, gin.H{
		"Removed": servicePlanet.RemoveByName(name),
	})
}
