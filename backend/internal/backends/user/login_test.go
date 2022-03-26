package user_test

import (
	"testing"
	"virtual-white-board-service/internal/backends/user"
	"virtual-white-board-service/internal/database"
	"virtual-white-board-service/internal/flags"
	"virtual-white-board-service/internal/models"
	"virtual-white-board-service/internal/services"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestLogin(db *gorm.DB) error {
	err := database.DropTables(db)
	if err != nil {
		return err
	}

	err = database.AutoMigrate(db)
	if err != nil {
		return err
	}

	tmp := &models.User{Username: "jbakke", Password: "password"}
	if result := db.Create(tmp); result.Error != nil {
		return result.Error
	}

	return nil
}

func TestLogin(t *testing.T) {
	flags, err := flags.ParseFlags()
	assert.NoError(t, err)

	db, err := database.New(flags)
	assert.NoError(t, err)

	err = setupTestLogin(db)
	assert.NoError(t, err)

	login := user.NewLogin(db, services.JWTAuthService())

	tests := map[string]struct {
		user      models.User
		expectErr bool
	}{
		"correct login": {user: models.User{Username: "jbakke", Password: "password"}},
		"wrong login":   {user: models.User{Username: "jbakke", Password: "wrong"}, expectErr: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			loginRes, err := login.Login(&test.user)
			if test.expectErr {
				assert.Error(t, err)
				assert.Nil(t, loginRes)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, loginRes)
				assert.Greater(t, len(loginRes.Token), 0)
			}
		})
	}
}
