# Cache-Control middleware for Gin

This Gin middleware generates cache-control headers.

[![Build Status](https://github.com/joeig/gin-cachecontrol/workflows/Tests/badge.svg)](https://github.com/joeig/gin-cachecontrol/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/joeig/gin-cachecontrol)](https://goreportcard.com/report/github.com/joeig/gin-cachecontrol)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/joeig/gin-cachecontrol)](https://pkg.go.dev/github.com/joeig/gin-cachecontrol)

## Usage

```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joeig/gin-cachecontrol/v2"
)

func main() {
	router := gin.Default()

	router.Use(cachecontrol.New(cachecontrol.Config{
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

	// Alternatively, you can choose a preset:
	// router.Use(cachecontrol.New(cachecontrol.NoCachePreset))

	router.GET("/", func(ginCtx *gin.Context) {
		ginCtx.String(http.StatusOK, "Hello, Gopher!")
	})

	router.Run()
}
```

## Documentation

See [GoDoc](https://godoc.org/github.com/joeig/gin-cachecontrol).
