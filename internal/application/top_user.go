package application

import (
	"context"
	"simple-redis-go/internal/domain"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func (s *AuthServer) GetTopUser(limit int) ([]domain.UserItemTopUser, error) {
	ctx := context.Background()

	users, err := s.rdb.ZRevRangeByScoreWithScores(ctx, CountUserPingKey, &redis.ZRangeBy{
		Offset: 0,
		Count:  int64(limit),
	}).Result()

	if err != nil {
		return nil, err
	}

	topUser := make([]domain.UserItemTopUser, len(users))
	for i, v := range users {
		var user domain.User
		idStr := v.Member.(string)
		id, ok := strconv.Atoi(idStr)
		if ok != nil {
			return nil, ok
		}
		if err := s.db.Find(&user, id).Error; err != nil {
			return nil, err
		}

		item := domain.UserItemTopUser{
			User:  user,
			Count: int64(v.Score),
		}
		topUser[i] = item
	}
	return topUser, nil
}
