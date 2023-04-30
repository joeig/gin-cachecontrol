# Cache-Control middleware for Gin

This Gin middleware generates cache-control headers.

[![Build Status](https://github.com/joeig/gin-cachecontrol/workflows/Tests/badge.svg)](https://github.com/joeig/gin-cachecontrol/actions)
[![Test coverage](https://img.shields.io/badge/coverage-100%25-success)](https://github.com/joeig/gin-cachecontrol/tree/master/.github/testcoverage.yml)
[![Go Report Card](https://goreportcard.com/badge/go.eigsys.de/gin-cachecontrol)](https://goreportcard.com/report/go.eigsys.de/gin-cachecontrol)
[![PkgGoDev](https://pkg.go.dev/badge/go.eigsys.de/gin-cachecontrol)](https://pkg.go.dev/go.eigsys.de/gin-cachecontrol)

## Usage

```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.eigsys.de/gin-cachecontrol/v2"
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

See [GoDoc](https://godoc.org/go.eigsys.de/gin-cachecontrol).
