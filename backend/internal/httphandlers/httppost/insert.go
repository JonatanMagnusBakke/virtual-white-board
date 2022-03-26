package httppost

import (
	"net/http"
	"virtual-white-board-service/internal/backends/post"
	"virtual-white-board-service/internal/models"

	"github.com/gin-gonic/gin"
)

//NewInsertHandler handler for post inserter
func NewInsertHandler(postInserter post.Inserter) func(c *gin.Context) {
	return func(c *gin.Context) {
		msg := new(models.Post)
		if err := c.ShouldBindJSON(msg); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err := postInserter.Insert(msg)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "")
	}
}
