package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pawoo-dev/pawoo-be/controller"
	"github.com/pawoo-dev/pawoo-be/dto"
)

func CreateServiceHandler(c *gin.Context) {
	var req dto.Service

	c.BindJSON(&req)
	user, _ := c.Get("user")
	req.CompanyId = user.(dto.User).CompanyId

	// todo: do validation

	err := controller.ServiceControllerObj.CreateService(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("success"))
	return
}

func GetServiceHandler(c *gin.Context) {
	companyId := c.Query("company_id")
	companyIdInt, err := strconv.Atoi(companyId)

	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	// todo: do validation

	serviceList, err := controller.ServiceControllerObj.GetService(companyIdInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, serviceList)
	return
}
