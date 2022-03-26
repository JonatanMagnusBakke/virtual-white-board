package models

//Post model
type Post struct {
	ID       int       `json:"id" gorm:"primary_key"`
	Message  string    `json:"message"`
	UserID   int       `json:"user_id" binding:"required"`
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
