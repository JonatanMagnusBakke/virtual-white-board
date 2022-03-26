package httpuser_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"virtual-white-board-service/internal/backends/user"
	"virtual-white-board-service/internal/dto"
	"virtual-white-board-service/internal/httphandlers"
	"virtual-white-board-service/internal/httphandlers/httpuser"
	"virtual-white-board-service/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type UserLoginMock struct {
	err error
}

func NewUserLoginMock(err error) user.Login {
	return &UserLoginMock{err: err}
}

func (m UserLoginMock) Login(user *models.User) (*dto.LoginResponse, error) {
	return nil, m.err
}

func TestInsert(t *testing.T) {
	tests := map[string]struct {
		input models.User
	}{
		"with username and password": {input: models.User{Username: "jbakke", Password: "password"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			login := NewUserLoginMock(nil)
			w := httptest.NewRecorder()
			g := gin.Default()
			g.POST(httphandlers.LoginRoute, httpuser.NewLoginHandler(login))

			buf, err := json.Marshal(test.input)
			assert.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, httphandlers.LoginRoute, bytes.NewReader(buf))

			// Test
			g.ServeHTTP(w, req)

			// Verify
			require.Equal(t, http.StatusOK, w.Code)
			require.Equal(t, "null", w.Body.String())
		})
	}
}
