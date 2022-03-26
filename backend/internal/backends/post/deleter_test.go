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

func setupTestDeleter(db *gorm.DB) error {
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

//TestDeleter test persisting post in the database and finding them
func TestDeleter(t *testing.T) {
	flags, err := flags.ParseFlags()
	require.NoError(t, err)

	db, err := database.New(flags)
	require.NoError(t, err)

	err = setupTestInserter(db)
	require.NoError(t, err)

	deleter := post.NewDeleter(db)

	tests := map[string]struct {
		post        models.Post
		expectError bool
	}{
		"default message": {post: models.Post{UserID: 1}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := db.Create(&test.post).Error
			require.NoError(t, err)

			err = deleter.Delete(&test.post)
			require.NoError(t, err)

			var dbPost models.Post
			db.First(dbPost)
			assert.Equal(t, 0, dbPost.ID)
		})
	}
}
