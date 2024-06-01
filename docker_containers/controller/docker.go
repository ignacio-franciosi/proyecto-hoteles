package controller

import (
	"docker_containers/client"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {

	stats, err := client.GetInfo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func GetInfoByService(c *gin.Context) {

	service := c.Param("service")

	stats, err := client.GetInfoByService(service)

	if err != nil {

		if err.Error() == "service does not exist" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func ScaleService(c *gin.Context) {

	service := c.Param("service")

	qty, err := client.ScaleService(service)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("Service %s scaled correctly to %d instances", service, qty)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func DeleteContainer(c *gin.Context) {

	id := c.Param("id")

	err := client.DeleteContainer(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("container %s deleted successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func GetScalableServices(c *gin.Context) {

	services := client.GetScalableServices()
	c.JSON(http.StatusOK, services)

}
