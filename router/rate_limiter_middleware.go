package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimitingMiddleware(period time.Duration, limit int) gin.HandlerFunc {

	rate := limiter.Rate{
		Period: period,
		Limit:  int64(limit),
	}

	store := memory.NewStore()
	// Alternatively, you can pass options to the limiter instance with several options.
	instance := limiter.New(store, rate)

	// Finally, give the limiter instance to your middleware initializer.

	return mgin.NewMiddleware(instance)
}
