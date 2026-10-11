package service

import (
	"context"
	"strings"
	"time"

	nexerrors "github.com/ycodeclub/Nex/errors"
	nexjwt "github.com/ycodeclub/nex-contrib/jwt"
	"github.com/ycodeclub/nex-demo/internal/domain"
	"github.com/ycodeclub/nex-demo/internal/repository"
)

// AuthService 登录与个人资料用例。
type AuthService interface {
	// Login 校验凭证并签发 JWT。
	Login(ctx context.Context, username, password string) (*LoginResult, error)
	// GetProfile 按用户 ID 返回资料。
	GetProfile(ctx context.Context, userID string) (*domain.User, error)
}

// LoginResult 登录成功结果。
type LoginResult struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
}

type authService struct {
	repo   repository.Repository
	tokens *nexjwt.Manager
	ttl    time.Duration
}

// NewAuthService 创建认证服务。
func NewAuthService(repo repository.Repository, tokens *nexjwt.Manager) AuthService {
	return &authService{
		repo:   repo,
		tokens: tokens,
		ttl:    24 * time.Hour,
	}
}

// Login 校验凭证并签发 JWT。
func (s *authService) Login(ctx context.Context, username, password string) (*LoginResult, error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	if username == "" || password == "" {
		return nil, nexerrors.ErrBadRequest.WithMessage("用户名和密码不能为空")
	}
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if !user.MatchPassword(password) {
		return nil, nexerrors.ErrUnauthorized.WithMessage("用户名或密码错误")
	}
	token, err := s.tokens.Generate(user.ID(), user.Roles(), map[string]string{
		"username": user.Username(),
	})
	if err != nil {
		return nil, nexerrors.ErrInternal.WithMessage("签发令牌失败")
	}
	return &LoginResult{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int64(s.ttl.Seconds()),
	}, nil
}

// GetProfile 按用户 ID 返回资料。
func (s *authService) GetProfile(ctx context.Context, userID string) (*domain.User, error) {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, nexerrors.ErrUnauthorized.WithMessage("未登录")
	}
	return s.repo.FindByID(ctx, userID)
}
