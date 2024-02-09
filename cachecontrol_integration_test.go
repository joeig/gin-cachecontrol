package cachecontrol_test

import (
	"context"
	cachecontrol "go.eigsys.de/gin-cachecontrol/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	testMethod = http.MethodGet
	testPath   = "/"
)

func newTestRouter(config cachecontrol.Config) *gin.Engine {
	router := gin.New()
	router.Use(cachecontrol.New(config))
	router.Handle(testMethod, testPath, func(ginCtx *gin.Context) {
		ginCtx.Data(http.StatusOK, "", []byte{})
	})

	return router
}

func requestTestRouter(handler http.Handler) *httptest.ResponseRecorder {
	request, _ := http.NewRequestWithContext(context.Background(), testMethod, testPath, nil)
	responseWriter := httptest.NewRecorder()

	handler.ServeHTTP(responseWriter, request)

	return responseWriter
}

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		config cachecontrol.Config
		want   string
	}{
		{
			name: "StandardCase",
			config: cachecontrol.Config{
				MustRevalidate: true,
				NoCache:        true,
				NoStore:        true,
				MaxAge:         cachecontrol.Duration(0),
			},
			want: "must-revalidate, no-cache, no-store, max-age=0",
		},
		{
			name:   "EmptyHeader",
			config: cachecontrol.Config{},
			want:   "",
		},
		{
			name: "RoundDuration",
			config: cachecontrol.Config{
				MaxAge:               cachecontrol.Duration(1500 * time.Millisecond),
				SMaxAge:              cachecontrol.Duration(1500 * time.Millisecond),
				StaleWhileRevalidate: cachecontrol.Duration(1500 * time.Millisecond),
				StaleIfError:         cachecontrol.Duration(1500 * time.Millisecond),
			},
			want: "max-age=2, s-maxage=2, stale-while-revalidate=2, stale-if-error=2",
		},
		{
			name: "AllFields",
			config: cachecontrol.Config{
				MustRevalidate:       true,
				NoCache:              true,
				NoStore:              true,
				NoTransform:          true,
				Public:               true,
				Private:              true,
				ProxyRevalidate:      true,
				MaxAge:               cachecontrol.Duration(1 * time.Second),
				SMaxAge:              cachecontrol.Duration(2 * time.Second),
				Immutable:            true,
				StaleWhileRevalidate: cachecontrol.Duration(3 * time.Second),
				StaleIfError:         cachecontrol.Duration(4 * time.Second),
			},
			want: "must-revalidate, no-cache, no-store, no-transform, public, private, proxy-revalidate, max-age=1, s-maxage=2, immutable, stale-while-revalidate=3, stale-if-error=4",
		},
		{
			name:   "NoCachePreset",
			config: cachecontrol.NoCachePreset,
			want:   "must-revalidate, no-cache, no-store",
		},
		{
			name:   "CacheAssetsForeverPreset",
			config: cachecontrol.CacheAssetsForeverPreset,
			want:   "public, max-age=31536000, immutable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newTestRouter(tt.config)
			response := requestTestRouter(router)

			if got := response.Header().Get(cachecontrol.CacheControlHeader); got != tt.want {
				t.Errorf("got=%q want=%q", got, tt.want)
			}
		})
	}
}
