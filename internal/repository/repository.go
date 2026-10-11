package repository

import (
	"context"
	"sync"

	nexerrors "github.com/ycodeclub/Nex/errors"
	"github.com/ycodeclub/nex-demo/internal/domain"
	"github.com/ycodeclub/nex-demo/internal/repository/dao"
)

// Repository 用户仓储。
type Repository interface {
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

type memoryStore struct {
	mu     sync.RWMutex
	byID   map[string]*domain.User
	byName map[string]*domain.User
}

// NewRepository 创建内存仓储，预置演示账号。
func NewRepository() Repository {
	s := &memoryStore{
		byID:   make(map[string]*domain.User),
		byName: make(map[string]*domain.User),
	}
	for _, row := range seedUserRows() {
		u := fromRow(row)
		s.byID[u.ID()] = u
		s.byName[u.Username()] = u
	}
	return s
}

// seedUserRows 返回演示用户行数据。
func seedUserRows() []dao.UserRow {
	return []dao.UserRow{
		{ID: "u-alice", Username: "alice", Password: "password123", DisplayName: "Alice", Roles: []string{"user", "admin"}},
		{ID: "u-bob", Username: "bob", Password: "password123", DisplayName: "Bob", Roles: []string{"user"}},
	}
}

// fromRow 将 dao 行转换为领域对象。
func fromRow(row dao.UserRow) *domain.User {
	u, err := domain.NewUser(row.ID, row.Username, row.Password, row.DisplayName, row.Roles)
	if err != nil {
		panic(err)
	}
	return u
}

// FindByUsername 按用户名查找用户。
func (s *memoryStore) FindByUsername(_ context.Context, username string) (*domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.byName[username]
	if !ok {
		return nil, nexerrors.ErrUnauthorized.WithMessage("用户名或密码错误")
	}
	return u, nil
}

// FindByID 按 ID 查找用户。
func (s *memoryStore) FindByID(_ context.Context, id string) (*domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.byID[id]
	if !ok {
		return nil, nexerrors.ErrNotFound.WithMessage("用户不存在")
	}
	return u, nil
}
