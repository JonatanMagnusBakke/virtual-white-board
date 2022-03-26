package comment_test

import (
	"testing"
	"virtual-white-board-service/internal/backends/comment"
	"virtual-white-board-service/internal/database"
	"virtual-white-board-service/internal/flags"
	"virtual-white-board-service/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestInserter(db *gorm.DB) error {
	err := database.DropTables(db)
	if err != nil {
		return err
	}

	err = database.AutoMigrate(db)
	if err != nil {
		return err
	}

	user := &models.User{Username: "jbakke", Password: "password"}

	err = db.Create(user).Error
	if err != nil {
		return err
	}

	post := &models.Post{Message: "message", UserID: user.ID}

	err = db.Create(post).Error
	if err != nil {
		return err
	}

	return nil
}

//TestInserter test persisting post in the database and finding them
func TestInserter(t *testing.T) {
	flags, err := flags.ParseFlags()
	require.NoError(t, err)

	db, err := database.New(flags)
	require.NoError(t, err)

	err = setupTestInserter(db)
	require.NoError(t, err)

	inserter := comment.NewInserter(db)

	tests := map[string]struct {
		comment     models.Comment
		expectError bool
	}{
		"with post id":    {comment: models.Comment{Message: "comment", PostID: 1}},
		"with no post id": {comment: models.Comment{Message: "comment"}, expectError: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			//Insert post
			err = inserter.Insert(&test.comment)
			if test.expectError {
				assert.Error(t, err)

			} else {
				assert.NoError(t, err)

				//Check post gets ID
				assert.NotEqual(t, test.comment.ID, 0)

				//Read post from database
				var comment models.Comment
				result := db.First(&comment, test.comment.ID)
				assert.NoError(t, result.Error)

				//Check that fields match
				assert.Equal(t, comment.Message, test.comment.Message)
			}

		})
	}
}
