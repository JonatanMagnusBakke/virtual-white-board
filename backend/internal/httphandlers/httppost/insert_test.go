package httppost_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"virtual-white-board-service/internal/backends/post"
	"virtual-white-board-service/internal/httphandlers"
	"virtual-white-board-service/internal/httphandlers/httppost"
	"virtual-white-board-service/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type PostInserterMock struct {
	err error
}

func NewPostInserterMock(err error) post.Inserter {
	return &PostInserterMock{err: err}
}

func (m PostInserterMock) Insert(post *models.Post) error {
	return m.err
}

func TestInsert(t *testing.T) {

	tests := map[string]struct {
		input              models.Post
		expectedStatusCode int
	}{
		"with User id":             {input: models.Post{UserID: 1}, expectedStatusCode: http.StatusOK},
		"with message and User id": {input: models.Post{UserID: 1, Message: "message"}, expectedStatusCode: http.StatusOK},
		"missing user id":          {input: models.Post{Message: "message"}, expectedStatusCode: http.StatusBadRequest},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			inserter := NewPostInserterMock(nil)
			w := httptest.NewRecorder()
			g := gin.Default()
			g.POST(httphandlers.PostsRoute, httppost.NewInsertHandler(inserter))

			buf, err := json.Marshal(test.input)
			assert.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, httphandlers.PostsRoute, bytes.NewReader(buf))

			// Test
			g.ServeHTTP(w, req)

			// Verify
			require.Equal(t, test.expectedStatusCode, w.Code)
		})
	}

}
