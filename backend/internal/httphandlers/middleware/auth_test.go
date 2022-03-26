package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"virtual-white-board-service/internal/httphandlers"
	"virtual-white-board-service/internal/httphandlers/middleware"
	"virtual-white-board-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		token              string
		expectedStatusCode int
	}{
		"emtpy token":      {token: "", expectedStatusCode: http.StatusUnauthorized},
		"with valid token": {token: services.JWTAuthService().GenerateToken("jbakke"), expectedStatusCode: http.StatusOK},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			g := gin.Default()

			g.Use(middleware.AuthMiddleware(services.JWTAuthService()))
			g.POST(httphandlers.UsersRoute, func(c *gin.Context) {
				c.String(http.StatusOK, "")
			})

			req := httptest.NewRequest(http.MethodPost, httphandlers.UsersRoute, nil)
			req.Header.Set("Authorization", test.token)
			// Test
			g.ServeHTTP(w, req)

			// Verify
			require.Equal(t, test.expectedStatusCode, w.Code)
		})
	}

}
