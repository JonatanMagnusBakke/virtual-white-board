package models

//Comment model
type Comment struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Message string `json:"message"`
	PostID  uint   `json:"post_id" binding:"required"`
}
