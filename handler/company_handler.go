package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pawoo-dev/pawoo-be/controller"
)

func GetAllCompany(c *gin.Context) {
	companyId := c.Query("company_id")
	if companyId != "" {
		companyIdInt, err := strconv.Atoi(companyId)
		if err != nil {
			c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
			return
		}
		details, err := controller.CompanyControllerObj.GetCompanyDetails(companyIdInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
			return
		}
		c.JSON(http.StatusOK, details)
		return
	}
	companyList, err := controller.CompanyControllerObj.GetAllCompany()
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, companyList)

}
