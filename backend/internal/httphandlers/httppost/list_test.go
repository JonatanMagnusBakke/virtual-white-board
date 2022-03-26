package httppost_test

import (
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

type PostListerMock struct {
	err   error
	posts *[]models.Post
}

func NewPostListerMock(err error, number int) post.Lister {
	var posts []models.Post
	for i := 0; i < number; i++ {
		posts = append(posts, models.Post{})
	}
	return &PostListerMock{err: err, posts: &posts}
}

func (m PostListerMock) List() (*[]models.Post, error) {
	return m.posts, m.err
}

func TestList(t *testing.T) {
	lister := NewPostListerMock(nil, 10)
	w := httptest.NewRecorder()
	g := gin.Default()
	g.GET(httphandlers.PostsRoute, httppost.NewListHandler(lister))

	req := httptest.NewRequest(http.MethodGet, httphandlers.PostsRoute, nil)

	// Test
	g.ServeHTTP(w, req)

	// Verify
	require.Equal(t, http.StatusOK, w.Code)

	var res *[]models.Post
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	assert.Equal(t, 10, len(*res))
}
