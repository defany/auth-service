package accessservice

import (
	"context"

	"github.com/defany/auth-service/app/internal/utils"
)

func (s *Service) DoesHaveAccess(ctx context.Context, accessToken string, endpoint string) error {
	claims, err := utils.VerifyToken(accessToken, s.jwtConfig.AccessSecretKey())
	if err != nil {
		return err
	}

	err = s.userRepo.DoesHaveAccess(ctx, claims.Role, endpoint)
	if err != nil {
		return err
	}

	return nil
}
