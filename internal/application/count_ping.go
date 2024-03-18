package application

import (
	"context"
	"fmt"
)

const CountUserPingKey = "ping_count"

func (s *AuthServer) AddPingToCount(ctx context.Context, id int64) error {
	key := fmt.Sprintf("%d", id)
	err := s.rdb.ZIncrBy(ctx, CountUserPingKey, 1, key).Err()
	return err
}

func (s *AuthServer) GetPingCount(ctx context.Context, id int64) (int, error) {
	key := fmt.Sprintf("%d", id)
	count, err := s.rdb.ZScore(ctx, CountUserPingKey, key).Result()
	return int(count), err
}
