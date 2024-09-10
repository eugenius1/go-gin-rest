package album

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	albumrepo "github.com/eugenius1/go-gin-rest/internal/db/repository/album"
	"github.com/eugenius1/go-gin-rest/internal/service/album"
)

func TestController_Integration(t *testing.T) {
	t.Parallel()

	repo := albumrepo.NewMemoryRepo()
	service := album.NewService(repo)
	c := &Controller{
		Service: service,
	}

	objID := "4"
	objJSON := fmt.Sprintf(`{"id":"%s","title":"Title","artist":"Artist","price":9.99}`, objID)

	t.Run("postAlbums", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		url, err := url.Parse("/albums")
		require.NoError(t, err)

		ctx.Request = &http.Request{
			Method: http.MethodPost,
			URL:    url,
			Header: http.Header{"Content-Type": []string{"application/json"}},
			Body:   io.NopCloser(strings.NewReader(objJSON)),
		}

		c.postAlbums(ctx)
		require.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, objJSON, w.Body.String())
	})

	t.Run("getAlbums", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		url, err := url.Parse("/albums")
		require.NoError(t, err)

		ctx.Request = &http.Request{
			Method: http.MethodGet,
			URL:    url,
		}

		c.getAlbums(ctx)
		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("getAlbumByID", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		url, err := url.Parse("/albums/" + objID)
		require.NoError(t, err)

		ctx.Request = &http.Request{
			Method: http.MethodGet,
			URL:    url,
		}
		ctx.AddParam("id", objID)

		t.Log("url=", url.String())

		c.getAlbumByID(ctx)
		require.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, objJSON, w.Body.String())
	})
}
