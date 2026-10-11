package transport

import (
	"context"

	nexerrors "github.com/ycodeclub/Nex/errors"
	"github.com/ycodeclub/Nex/middleware/auth"
	v1 "github.com/ycodeclub/nex-demo/proto/v1"
)

// Login 实现 AuthServer。
func (h *Handler) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	if req == nil {
		return nil, nexerrors.ErrBadRequest.WithMessage("请求体不能为空")
	}
	result, err := h.svc.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		AccessToken: result.AccessToken,
		TokenType:   result.TokenType,
		ExpiresIn:   result.ExpiresIn,
	}, nil
}

// GetProfile 实现 AuthServer。
func (h *Handler) GetProfile(ctx context.Context, _ *v1.GetProfileRequest) (*v1.GetProfileReply, error) {
	userID, ok := auth.UserIDFromContext(ctx)
	if !ok || userID == "" {
		return nil, nexerrors.ErrUnauthorized.WithMessage("未登录")
	}
	user, err := h.svc.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileReply{
		UserId:      user.ID(),
		Username:    user.Username(),
		DisplayName: user.DisplayName(),
		Roles:       user.Roles(),
	}, nil
}
