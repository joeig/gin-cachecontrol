package cachecontrol

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const CacheControlHeader = "Cache-Control"

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

// ConfigOption mutates a cache-control configuration.
type ConfigOption func(*Config)

// WithMustRevalidate configures the must-revalidate directive.
func WithMustRevalidate(value bool) ConfigOption {
	return func(config *Config) {
		config.MustRevalidate = value
	}
}

// WithNoCache configures the no-cache directive.
func WithNoCache(value bool) ConfigOption {
	return func(config *Config) {
		config.NoCache = value
	}
}

// WithNoStore configures the no-store directive.
func WithNoStore(value bool) ConfigOption {
	return func(config *Config) {
		config.NoStore = value
	}
}

// WithNoTransform configures the no-transform directive.
func WithNoTransform(value bool) ConfigOption {
	return func(config *Config) {
		config.NoTransform = value
	}
}

// WithPublic configures the public directive.
func WithPublic(value bool) ConfigOption {
	return func(config *Config) {
		config.Public = value
	}
}

// WithPrivate configures the private directive.
func WithPrivate(value bool) ConfigOption {
	return func(config *Config) {
		config.Private = value
	}
}

// WithProxyRevalidate configures the proxy-revalidate directive.
func WithProxyRevalidate(value bool) ConfigOption {
	return func(config *Config) {
		config.ProxyRevalidate = value
	}
}

// WithMaxAge configures the max-age directive.
func WithMaxAge(value *time.Duration) ConfigOption {
	return func(config *Config) {
		config.MaxAge = value
	}
}

// WithSMaxAge configures the s-maxage directive.
func WithSMaxAge(value *time.Duration) ConfigOption {
	return func(config *Config) {
		config.SMaxAge = value
	}
}

// WithImmutable configures the immutable directive.
func WithImmutable(value bool) ConfigOption {
	return func(config *Config) {
		config.Immutable = value
	}
}

// WithStaleWhileRevalidate configures the stale-while-revalidate directive.
func WithStaleWhileRevalidate(value *time.Duration) ConfigOption {
	return func(config *Config) {
		config.StaleWhileRevalidate = value
	}
}

// WithStaleIfError configures the stale-if-error directive.
func WithStaleIfError(value *time.Duration) ConfigOption {
	return func(config *Config) {
		config.StaleIfError = value
	}
}

func (c *Config) buildCacheControl() string {
	var values []string

	if c.MustRevalidate {
		values = append(values, "must-revalidate")
	}

	if c.NoCache {
		values = append(values, "no-cache")
	}

	if c.NoStore {
		values = append(values, "no-store")
	}

	if c.NoTransform {
		values = append(values, "no-transform")
	}

	if c.Public {
		values = append(values, "public")
	}

	if c.Private {
		values = append(values, "private")
	}

	if c.ProxyRevalidate {
		values = append(values, "proxy-revalidate")
	}

	if c.MaxAge != nil {
		values = append(values, fmt.Sprintf("max-age=%.f", c.MaxAge.Seconds()))
	}

	if c.SMaxAge != nil {
		values = append(values, fmt.Sprintf("s-maxage=%.f", c.SMaxAge.Seconds()))
	}

	if c.Immutable {
		values = append(values, "immutable")
	}

	if c.StaleWhileRevalidate != nil {
		values = append(values, fmt.Sprintf("stale-while-revalidate=%.f", c.StaleWhileRevalidate.Seconds()))
	}

	if c.StaleIfError != nil {
		values = append(values, fmt.Sprintf("stale-if-error=%.f", c.StaleIfError.Seconds()))
	}

	return strings.Join(values, ", ")
}

func (c *Config) apply(ginCtx *gin.Context, value string) {
	header := ginCtx.Writer.Header()
	header.Set(CacheControlHeader, value)
}

// New creates a new Gin middleware which generates a cache-control header.
// Existing cache-control headers are removed.
// Other caching-related headers, such as `Expires` and `Pragma`, remain unchanged.
func New(config Config) gin.HandlerFunc {
	value := config.buildCacheControl()

	return func(ginCtx *gin.Context) {
		config.apply(ginCtx, value)
	}
}

// NewWithOptions creates a new Gin middleware which generates a cache-control header
// from functional options. Existing cache-control headers are removed.
// Other caching-related headers, such as `Expires` and `Pragma`, remain unchanged.
func NewWithOptions(options ...ConfigOption) gin.HandlerFunc {
	config := ApplyOptionsToConfig(Config{}, options...)
	return New(config)
}

// ApplyOptionsToConfig applies functional options to a cache-control configuration.
// Nil options are ignored.
func ApplyOptionsToConfig(config Config, options ...ConfigOption) Config {
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&config)
	}
	return config
}

// Duration is a helper function which returns a time.Duration pointer.
func Duration(duration time.Duration) *time.Duration {
	return &duration
}

// NoCachePreset is a cache-control configuration preset which advices the HTTP client not to cache at all.
var NoCachePreset = Config{
	MustRevalidate: true,
	NoCache:        true,
	NoStore:        true,
}

// CacheAssetsForeverPreset is a cache-control configuration preset which advices the HTTP client
// and all caches in between to cache the object forever without revalidation.
// Technically, "forever" means 1 year, in order to comply with common CDN limits.
var CacheAssetsForeverPreset = Config{
	Public:    true,
	MaxAge:    Duration(8760 * time.Hour),
	Immutable: true,
}
