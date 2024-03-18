package application

import (
	"context"
	"time"
)

const PingLockKey = "ping_lock"

func (s *AuthServer) Ping(session string) (int, error) {
	ctx := context.Background()
	rs, err := s.rdb.SetNX(ctx, PingLockKey, true, 0).Result()
	defer s.rdb.Del(ctx, PingLockKey)

	if !rs {
		return 0, ErrPing
	}
	if err != nil {
		return 0, ErrDBInternal
	}
	user, err := s.GetUserFromSession(session)

	if err != nil {
		return 0, err
	}

	id := int64(user.ID)
	err = s.RateLimit(ctx, id)
	if err != nil {
		return 0, err
	}

	s.AddPingToCount(ctx, id)

	s.AddPingCountUser(id)
	count, _ := s.GetPingCount(ctx, id)

	defer time.Sleep(5 * time.Second)

	return count, nil
}
