package post_test

import (
	"testing"
	"virtual-white-board-service/internal/backends/post"
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

	inserter := post.NewInserter(db)

	tests := map[string]struct {
		post        models.Post
		expectError bool
	}{
		"default message": {post: models.Post{UserID: 1}},
		"with message":    {post: models.Post{Message: "message", UserID: 1}},
		"with no userID":  {post: models.Post{Message: "message"}, expectError: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			//Insert post
			err = inserter.Insert(&test.post)
			if test.expectError {
				assert.Error(t, err)

			} else {
				assert.NoError(t, err)

				//Check post gets ID
				assert.NotEqual(t, test.post.ID, 0)

				//Read post from database
				var post models.Post
				result := db.First(&post, test.post.ID)
				assert.NoError(t, result.Error)

				//Check that fields match
				assert.Equal(t, post.Message, test.post.Message)
			}

		})
	}
}
