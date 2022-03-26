package post

import (
	"virtual-white-board-service/internal/models"

	"gorm.io/gorm"
)

//Lister interface for listing posts
type Lister interface {
	List() (*[]models.Post, error)
}

type lister struct {
	db *gorm.DB
}

//NewLister creates new post lister
func NewLister(db *gorm.DB) Lister {
	res := lister{db: db}
	return res
}

//Insert persists the post to the database
func (m lister) List() (*[]models.Post, error) {
	var posts *[]models.Post
	if result := m.db.Preload("Comments").Preload("User").Find(&posts).Error; result != nil {
		return nil, result
	}
	return posts, nil
}
