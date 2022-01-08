package cachecontrol

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const cacheControlHeader = "Cache-Control"

// Config defines a cache-control configuration.
//
// References:
// https://datatracker.ietf.org/doc/html/rfc7234#section-5.2.2
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
type Config struct {
	MustRevalidate       bool
	NoCache              bool
	NoStore              bool
	NoTransform          bool
	Public               bool
	Private              bool
	ProxyRevalidate      bool
	MaxAge               *time.Duration
	SMaxAge              *time.Duration
	Immutable            bool
	StaleWhileRevalidate *time.Duration
	StaleIfError         *time.Duration
}

// New creates a new Gin middleware which generates a cache-control header.
func New(config *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.apply(c)
	}
}

func (c *Config) buildCacheControl() string {
	var cc []string

	if c.MustRevalidate {
		cc = append(cc, "must-revalidate")
	}

	if c.NoCache {
		cc = append(cc, "no-cache")
	}

	if c.NoStore {
		cc = append(cc, "no-store")
	}

	if c.NoTransform {
		cc = append(cc, "no-transform")
	}

	if c.Public {
		cc = append(cc, "public")
	}

	if c.Private {
		cc = append(cc, "private")
	}

	if c.ProxyRevalidate {
		cc = append(cc, "proxy-revalidate")
	}

	if c.MaxAge != nil {
		cc = append(cc, fmt.Sprintf("max-age=%.f", c.MaxAge.Seconds()))
	}

	if c.SMaxAge != nil {
		cc = append(cc, fmt.Sprintf("s-maxage=%.f", c.SMaxAge.Seconds()))
	}

	if c.Immutable {
		cc = append(cc, "immutable")
	}

	if c.StaleWhileRevalidate != nil {
		cc = append(cc, fmt.Sprintf("stale-while-revalidate=%.f", c.StaleWhileRevalidate.Seconds()))
	}

	if c.StaleIfError != nil {
		cc = append(cc, fmt.Sprintf("stale-if-error=%.f", c.StaleIfError.Seconds()))
	}

	return strings.Join(cc, ", ")
}

func (c *Config) apply(ginCtx *gin.Context) {
	header := ginCtx.Writer.Header()
	header.Set(cacheControlHeader, c.buildCacheControl())
}

// NoCachePreset is a cache-control configuration preset which advices the HTTP client not to cache at all.
var NoCachePreset = &Config{
	MustRevalidate: true,
	NoCache:        true,
	NoStore:        true,
}

// CacheAssetsForeverPreset is a cache-control configuration preset which advices the HTTP client
// and all caches in between to cache the object forever without revalidation.
// Technically, "forever" means 1 year, in order to comply with common CDN limits.
var CacheAssetsForeverPreset = &Config{
	Public:    true,
	MaxAge:    Duration(8760 * time.Hour),
	Immutable: true,
}

// Duration is a helper function which returns a time.Duration pointer.
func Duration(duration time.Duration) *time.Duration {
	return &duration
}
