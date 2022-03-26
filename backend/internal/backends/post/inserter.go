package post

import (
	"virtual-white-board-service/internal/models"

	"gorm.io/gorm"
)

//Inserter interface for inserting posts
type Inserter interface {
	Insert(post *models.Post) error
}

type inserter struct {
	db *gorm.DB
}

//NewInserter creates new post inserter
func NewInserter(db *gorm.DB) Inserter {
	res := inserter{db: db}
	return res
}

//Insert persists the post to the database
func (m inserter) Insert(post *models.Post) error {
	if result := m.db.Create(post); result.Error != nil {
		return result.Error
	}
	return nil
}
