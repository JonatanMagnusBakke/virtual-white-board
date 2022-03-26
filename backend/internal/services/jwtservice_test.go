package services_test

import (
	"testing"
	"virtual-white-board-service/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestJWTServiceGenerateAndValidate(t *testing.T) {
	jwtService := services.JWTAuthService()

	token := jwtService.GenerateToken("jbakke")

	assert.Greater(t, len(token), 0)

	res, err := jwtService.ValidateToken(token)
	assert.NoError(t, err)
	assert.True(t, res.Valid)
}

func TestJWTServiceValidateError(t *testing.T) {
	jwtService := services.JWTAuthService()
	_, err := jwtService.ValidateToken("")
	assert.Error(t, err)
}
