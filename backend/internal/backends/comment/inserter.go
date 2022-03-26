package comment

import (
	"virtual-white-board-service/internal/models"

	"gorm.io/gorm"
)

//Inserter interface for inserting posts
type Inserter interface {
	Insert(comment *models.Comment) error
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
func (m inserter) Insert(comment *models.Comment) error {
	if result := m.db.Create(comment); result.Error != nil {
		return result.Error
	}
	return nil
}
