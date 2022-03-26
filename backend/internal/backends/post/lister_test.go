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

func setupTestLister(db *gorm.DB, number int) error {
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

	for i := 0; i < number; i++ {
		err = db.Create(&models.Post{UserID: user.ID, Message: "message"}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

//TestLister test listing post from the database
func TestLister(t *testing.T) {
	flags, err := flags.ParseFlags()
	require.NoError(t, err)

	db, err := database.New(flags)
	require.NoError(t, err)

	err = setupTestLister(db, 10)
	require.NoError(t, err)

	lister := post.NewLister(db)

	posts, err := lister.List()
	require.NoError(t, err)
	assert.Equal(t, 10, len(*posts))
}
