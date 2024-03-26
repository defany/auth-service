package converter

import (
	"github.com/defany/auth-service/app/internal/model"
	authv1 "github.com/defany/auth-service/app/pkg/gen/proto/auth/v1"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/platcom/pkg/cond"
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
		Name:      input.Nickname,
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

func ToUserLogin(input *authv1.LoginRequest) model.UserLoginInput {
	return model.UserLoginInput{
		Nickname: input.GetUsername(),
		Password: input.GetPassword(),
	}
}

func FromUserLogin(input model.UserLoginOutput) *authv1.LoginResponse {
	return &authv1.LoginResponse{
		RefreshToken: input.RefreshToken,
	}
}
