package httpuser

import (
	"net/http"
	"virtual-white-board-service/internal/backends/user"
	"virtual-white-board-service/internal/models"

	"github.com/gin-gonic/gin"
)

//NewLoginHandler handler for user login
func NewLoginHandler(login user.Login) func(c *gin.Context) {
	return func(c *gin.Context) {
		msg := new(models.User)
		if err := c.ShouldBindJSON(msg); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		loginRes, err := login.Login(msg)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, loginRes)
	}
}
