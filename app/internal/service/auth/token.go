package authservice

import (
	"context"

	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/auth-service/app/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, input model.UserLoginInput) (model.UserLoginOutput, error) {
	user, err := s.userRepo.UserByNickname(ctx, input.Nickname)
	if err != nil {
		return model.UserLoginOutput{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return model.UserLoginOutput{}, err
	}

	jwtConfig := s.jwtConfig

	refreshToken, err := utils.GenerateToken(user, jwtConfig.RefreshSecretKey(), jwtConfig.RefreshTokenDuration())
	if err != nil {
		return model.UserLoginOutput{}, err
	}

	return model.UserLoginOutput{
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	jwtConfig := s.jwtConfig

	claims, err := utils.VerifyToken(refreshToken, jwtConfig.RefreshSecretKey())
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.UserByNickname(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	generatedRefreshToken, err := utils.GenerateToken(user, jwtConfig.RefreshSecretKey(), jwtConfig.RefreshTokenDuration())
	if err != nil {
		return "", err
	}

	return generatedRefreshToken, nil
}

func (s *Service) AccessToken(ctx context.Context, refreshToken string) (string, error) {
	jwtConfig := s.jwtConfig

	claims, err := utils.VerifyToken(refreshToken, jwtConfig.RefreshSecretKey())
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.UserByNickname(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	generatedAccessToken, err := utils.GenerateToken(user, jwtConfig.AccessSecretKey(), jwtConfig.AccessTokenDuration())
	if err != nil {
		return "", err
	}

	return generatedAccessToken, nil
}
