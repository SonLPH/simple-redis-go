package application

import (
	"context"
	"fmt"
)

const CountPing = "ping_count_user"

func (s *AuthServer) AddPingCountUser(id int64) error {
	ctx := context.Background()
	key := fmt.Sprintf("%d", id)
	err := s.rdb.PFAdd(ctx, CountPing, key).Err()
	return err
}

func (s *AuthServer) GetCountUser() (int64, error) {
	ctx := context.Background()
	count, err := s.rdb.PFCount(ctx, CountPing).Result()
	return count, err
}
