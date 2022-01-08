package cacheControl_test

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cacheControl "github.com/joeig/gin-cachecontrol"
)

func ExampleNew() {
	router := gin.Default()

	router.Use(cacheControl.New(&cacheControl.Config{
		MustRevalidate:       true,
		NoCache:              false,
		NoStore:              false,
		NoTransform:          false,
		Public:               true,
		Private:              false,
		ProxyRevalidate:      true,
		MaxAge:               cacheControl.Duration(30 * time.Minute),
		SMaxAge:              nil,
		Immutable:            false,
		StaleWhileRevalidate: cacheControl.Duration(2 * time.Hour),
		StaleIfError:         cacheControl.Duration(2 * time.Hour),
	}))

	router.GET("/", func(ginCtx *gin.Context) {
		ginCtx.String(http.StatusOK, "Hello, Gopher!")
	})

	router.Run()
}
