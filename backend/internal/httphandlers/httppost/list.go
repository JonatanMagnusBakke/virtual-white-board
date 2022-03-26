package httppost

import (
	"net/http"
	"virtual-white-board-service/internal/backends/post"
	"virtual-white-board-service/internal/models"

	"github.com/gin-gonic/gin"
)

//PostDTO dto
type PostDTO struct {
	ID       int              `json:"id" gorm:"primary_key"`
	Message  string           `json:"message"`
	UserID   int              `json:"user_id"`
	Comments []models.Comment `json:"comments"`
}

//NewListHandler handler for post lister
func NewListHandler(postLister post.Lister) func(c *gin.Context) {
	return func(c *gin.Context) {
		posts, err := postLister.List()
		var postsDTO []PostDTO
		for _, val := range *posts {
			postsDTO = append(postsDTO, PostDTO{ID: val.ID, Message: val.Message, Comments: val.Comments, UserID: val.UserID})
		}
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, postsDTO)
	}
}
