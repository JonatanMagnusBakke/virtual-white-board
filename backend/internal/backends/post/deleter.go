package post

import (
	"virtual-white-board-service/internal/models"

	"gorm.io/gorm"
)

//Deleter interface for deleting posts
type Deleter interface {
	Delete(post *models.Post) error
}

type deleter struct {
	db *gorm.DB
}

//NewDeleter creates new post inserter
func NewDeleter(db *gorm.DB) Deleter {
	res := deleter{db: db}
	return res
}

//Insert persists the post to the database
func (m deleter) Delete(post *models.Post) error {
	err := m.db.Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}
