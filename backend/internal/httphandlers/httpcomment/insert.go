package httpcomment

import (
	"net/http"
	"virtual-white-board-service/internal/backends/comment"
	"virtual-white-board-service/internal/models"

	"github.com/gin-gonic/gin"
)

//NewInsertHandler handler for post inserter
func NewInsertHandler(commentInserter comment.Inserter) func(c *gin.Context) {
	return func(c *gin.Context) {
		msg := new(models.Comment)
		if err := c.ShouldBindJSON(msg); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err := commentInserter.Insert(msg)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "")
	}
}
