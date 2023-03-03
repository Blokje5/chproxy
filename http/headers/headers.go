package headers

import (
	"fmt"
	"time"
)

// headers
const (
	// Content related headers
	ContentEncoding = "Content-Encoding"
	ContetType = "Content-Type"

	// Cache related headers
	XCache       = "X-Cache"
	CacheControl = "Cache-Control"
)

// values
const (
	// Content related headers
	ContentEncodingGZIP = "gzip"

	// Cache related values
	XCacheHit  = "HIT"
	XCacheMiss = "MISS"
)

func CacheControValue(ttl time.Duration) string {
	expireSeconds := uint(ttl / time.Second)
	return fmt.Sprintf("max-age=%d", expireSeconds)
}
