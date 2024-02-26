package converter

import (
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/auth-service/app/pkg/cond"
	userv1 "github.com/defany/auth-service/app/pkg/gen/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserCreate(input *userv1.CreateRequest) model.UserCreate {
	return model.UserCreate{
		Name:            input.GetName(),
		Email:           input.GetEmail(),
		Password:        input.GetPassword(),
		PasswordConfirm: input.GetPasswordConfirm(),
		Role:            input.GetRole().String(),
	}
}

func ToGetResponse(input model.User) *userv1.GetResponse {
	return &userv1.GetResponse{
		Id:        int64(input.ID),
		Name:      input.Name,
		Email:     input.Email,
		Role:      ToUserProtoRole(input.Role),
		CreatedAt: timestamppb.New(input.CreatedAt),
		UpdatedAt: cond.Ternary(!input.UpdatedAt.IsZero(), timestamppb.New(input.UpdatedAt), nil),
	}
}

func ToUserUpdate(input *userv1.UpdateRequest) model.UserUpdate {
	name := input.GetName()
	email := input.GetEmail()
	role := input.GetRole().String()

	return model.UserUpdate{
		ID:    uint64(input.GetId()),
		Name:  &name,
		Email: &email,
		Role:  &role,
	}
}

func ToUserProtoRole(role string) userv1.UserRole {
	return userv1.UserRole(userv1.UserRole_value[role])
}
