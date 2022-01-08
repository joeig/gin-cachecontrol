package cachecontrol_test

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joeig/gin-cachecontrol"
)

func ExampleNew() {
	router := gin.Default()

	router.Use(cachecontrol.New(&cachecontrol.Config{
		MustRevalidate:       true,
		NoCache:              false,
		NoStore:              false,
		NoTransform:          false,
		Public:               true,
		Private:              false,
		ProxyRevalidate:      true,
		MaxAge:               cachecontrol.Duration(30 * time.Minute),
		SMaxAge:              nil,
		Immutable:            false,
		StaleWhileRevalidate: cachecontrol.Duration(2 * time.Hour),
		StaleIfError:         cachecontrol.Duration(2 * time.Hour),
	}))

	router.GET("/", func(ginCtx *gin.Context) {
		ginCtx.String(http.StatusOK, "Hello, Gopher!")
	})

	router.Run()
}
