package application

import (
	"context"
	"fmt"
	"time"
)

const RateLimitKey = "ping_rate_limit:"

func (s *AuthServer) RateLimit(ctx context.Context, id int64) error {
	key := fmt.Sprintf("%s:%d", RateLimitKey, id)
	count, err := s.rdb.Incr(ctx, key).Result()
	if err != nil {
		return ErrDBInternal
	}

	if count > 2 {
		return ErrRateLimit
	}

	if count == 1 {
		s.rdb.ExpireNX(ctx, key, 2*time.Minute)
	}

	return nil
}
