package httpcomment_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"virtual-white-board-service/internal/backends/comment"
	"virtual-white-board-service/internal/httphandlers"
	"virtual-white-board-service/internal/httphandlers/httpcomment"
	"virtual-white-board-service/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CommentInserterMock struct {
	err error
}

func NewCommentInserterMock(err error) comment.Inserter {
	return &CommentInserterMock{err: err}
}

func (m CommentInserterMock) Insert(Comment *models.Comment) error {
	return m.err
}

func TestInsert(t *testing.T) {

	tests := map[string]struct {
		input              models.Comment
		expectedStatusCode int
	}{
		"with PostID":             {input: models.Comment{PostID: 1}, expectedStatusCode: http.StatusOK},
		"with message and PostID": {input: models.Comment{Message: "message", PostID: 1}, expectedStatusCode: http.StatusOK},
		"missing PostID":          {input: models.Comment{Message: "message"}, expectedStatusCode: http.StatusBadRequest},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			inserter := NewCommentInserterMock(nil)
			w := httptest.NewRecorder()
			g := gin.Default()
			g.POST(httphandlers.CommentRoute, httpcomment.NewInsertHandler(inserter))

			buf, err := json.Marshal(test.input)
			assert.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, httphandlers.CommentRoute, bytes.NewReader(buf))

			// Test
			g.ServeHTTP(w, req)

			// Verify
			require.Equal(t, test.expectedStatusCode, w.Code)

		})
	}

}
