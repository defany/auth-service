package auth

import (
	"context"

	"github.com/defany/auth-service/app/internal/converter"
	authv1 "github.com/defany/auth-service/app/pkg/gen/proto/auth/v1"
)

func (i *Implementation) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	output, err := i.service.Login(ctx, converter.ToUserLogin(req))
	if err != nil {
		return nil, err
	}

	return converter.FromUserLogin(output), nil
}

func (i *Implementation) GetRefreshToken(ctx context.Context, req *authv1.GetRefreshTokenRequest) (*authv1.GetRefreshTokenResponse, error) {
	generatedRefreshToken, err := i.service.RefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &authv1.GetRefreshTokenResponse{
		RefreshToken: generatedRefreshToken,
	}, nil
}

func (i *Implementation) GetAccessToken(ctx context.Context, req *authv1.GetAccessTokenRequest) (*authv1.GetAccessTokenResponse, error) {
	generatedAccessToken, err := i.service.AccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &authv1.GetAccessTokenResponse{
		AccessToken: generatedAccessToken,
	}, nil
}
