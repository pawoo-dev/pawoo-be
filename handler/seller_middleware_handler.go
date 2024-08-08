package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawoo-dev/pawoo-be/controller"
)

func CheckUserIsSellerMiddleware(c *gin.Context) {
	email, _ := c.Get("email")
	user, err := controller.UserControllerObj.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.Set("user", user)
	return
}
