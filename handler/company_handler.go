package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawoo-dev/pawoo-be/controller"
)

func GetAllCompany(c *gin.Context) {
	companyList, err := controller.CompanyControllerObj.GetAllCompany()
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, companyList)
	return
}
