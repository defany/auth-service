package converter

import (
	"github.com/defany/auth-service/app/internal/model"
	storemodel "github.com/defany/auth-service/app/internal/repository/user/model"
	"time"
)

func UserToModel(user storemodel.User) model.User {
	var updatedAt *time.Time

	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.V
	}

	return model.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            user.Role,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       updatedAt,
	}
}
