package cachecontrol

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

const testPath = "/"

func newTestRouter(config *Config) *gin.Engine {
	router := gin.New()
	router.Use(New(config))
	router.GET(testPath, func(ginCtx *gin.Context) {
		ginCtx.Data(http.StatusOK, "", []byte{})
	})

	return router
}

func requestTestRouter(handler http.Handler) *httptest.ResponseRecorder {
	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, testPath, nil)
	responseWriter := httptest.NewRecorder()

	handler.ServeHTTP(responseWriter, request)

	return responseWriter
}

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
		want   string
	}{
		{
			name: "StandardCase",
			config: &Config{
				MustRevalidate: true,
				NoCache:        true,
				NoStore:        true,
				MaxAge:         Duration(0),
			},
			want: "must-revalidate, no-cache, no-store, max-age=0",
		},
		{
			name:   "EmptyHeader",
			config: &Config{},
			want:   "",
		},
		{
			name:   "RoundDuration",
			config: &Config{MaxAge: Duration(1500 * time.Millisecond)},
			want:   "max-age=2",
		},
		{
			name: "AllFields",
			config: &Config{
				MustRevalidate:       true,
				NoCache:              true,
				NoStore:              true,
				NoTransform:          true,
				Public:               true,
				Private:              true,
				ProxyRevalidate:      true,
				MaxAge:               Duration(1 * time.Second),
				SMaxAge:              Duration(2 * time.Second),
				Immutable:            true,
				StaleWhileRevalidate: Duration(3 * time.Second),
				StaleIfError:         Duration(4 * time.Second),
			},
			want: "must-revalidate, no-cache, no-store, no-transform, public, private, proxy-revalidate, max-age=1, s-maxage=2, immutable, stale-while-revalidate=3, stale-if-error=4",
		},
		{
			name:   "NoCachePreset",
			config: NoCachePreset,
			want:   "must-revalidate, no-cache, no-store",
		},
		{
			name:   "CacheAssetsForeverPreset",
			config: CacheAssetsForeverPreset,
			want:   "public, max-age=31536000, immutable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newTestRouter(tt.config)
			response := requestTestRouter(router)

			if got := response.Header().Get(cacheControlHeader); got != tt.want {
				t.Errorf("got=%q want=%q", got, tt.want)
			}
		})
	}
}
