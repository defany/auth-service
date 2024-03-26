package converter

import (
	"github.com/defany/auth-service/app/internal/model"
	storemodel "github.com/defany/auth-service/app/internal/repository/user/model"
)

func UserToModel(user storemodel.User) model.User {
	return model.User{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt.V,
	}
}
