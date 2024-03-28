# Cache-Control middleware for Gin

This Gin middleware generates cache-control headers.

[![Build Status](https://github.com/joeig/gin-cachecontrol/actions/workflows/tests.yml/badge.svg)](https://github.com/joeig/gin-cachecontrol/actions/workflows/tests.yml)
[![Test coverage](https://img.shields.io/badge/coverage-100%25-success)](https://github.com/joeig/gin-cachecontrol/tree/master/.github/testcoverage.yml)
[![Go Report Card](https://goreportcard.com/badge/go.eigsys.de/gin-cachecontrol/v2)](https://goreportcard.com/report/go.eigsys.de/gin-cachecontrol/v2)
[![PkgGoDev](https://pkg.go.dev/badge/go.eigsys.de/gin-cachecontrol/v2)](https://pkg.go.dev/go.eigsys.de/gin-cachecontrol/v2)

## Setup

```shell
go get -u go.eigsys.de/gin-cachecontrol/v2
```

```go
import "go.eigsys.de/gin-cachecontrol/v2"
```

## Usage

### With a preset

```go
// Apply globally:
r.Use(cachecontrol.New(cachecontrol.NoCachePreset))

// Apply to specific routes:
cacheForever := cachecontrol.New(cachecontrol.CacheAssetsForeverPreset)
r.GET("/favicon.ico", cacheForever, faviconHandler)
```

Supported presets ([documentation](https://pkg.go.dev/go.eigsys.de/gin-cachecontrol/v2#pkg-variables)):

* `cachecontrol.NoCachePreset`
* `cachecontrol.CacheAssetsForeverPreset` (you may only want this for carefully selected routes)

### With a custom configuration

```go
r.Use(
    cachecontrol.New(
        cachecontrol.Config{
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
        }
    )
)
```

## Documentation

See [Go reference](https://pkg.go.dev/go.eigsys.de/gin-cachecontrol/v2).
