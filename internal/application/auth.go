package application

import (
	"context"
	"errors"
	"fmt"
	"simple-redis-go/internal/domain"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthServer struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewAuthServer(db *gorm.DB, rdb *redis.Client) *AuthServer {
	return &AuthServer{
		db:  db,
		rdb: rdb,
	}
}

func (s *AuthServer) CreateUser(email, password, firstName, lastName string) (*domain.User, error) {
	user := domain.NewUser(email, password, firstName, lastName)

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServer) FindUser(email string) (*domain.User, error) {
	var user *domain.User
	err := s.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, ErrDBInternal
	}
	return user, err
}

func (s *AuthServer) Login(email, password string) (sessionId string, err error) {
	var user *domain.User
	user, err = s.FindUser(email)

	if err != nil {
		return "", err
	}

	isVerified := user.VerifyPassword(password)
	if !isVerified {
		return "", ErrAuth
	}

	key := fmt.Sprintf("session_%d", time.Now().UnixNano())

	s.rdb.Set(context.Background(), key, *user, 2*time.Hour)
	return key, nil
}

func (s *AuthServer) GetUserFromSession(session string) (*domain.User, error) {
	userData, ok := s.db.Get("session_" + session)
	if !ok {
		return nil, ErrSessionInvalid
	}

	user := userData.(domain.User)
	return &user, nil
}
