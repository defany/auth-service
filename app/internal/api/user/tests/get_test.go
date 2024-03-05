package usertests

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/api/user"
	"github.com/defany/auth-service/app/internal/model"
	mockdefserv "github.com/defany/auth-service/app/internal/service/mocks"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/auth-service/app/pkg/logger/handlers/slogpretty"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"testing"
)

func TestImplementation_Get(t *testing.T) {
	type mocker func() *mockdefserv.MockUserService

	type args struct {
		ctx context.Context
		req *userv1.GetRequest
	}

	var (
		ctx = context.Background()

		serviceErr = status.Error(codes.Internal, "failed to get user by id")

		id = gofakeit.Uint64()

		req = &userv1.GetRequest{
			Id: int64(id),
		}

		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()
		role      = userv1.UserRole(int32(gofakeit.Float64Range(0, 2)))

		userModel = model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      role.String(),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		res = &userv1.GetResponse{
			Id:        int64(id),
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		}
	)

	tests := []struct {
		name string
		args args
		want *userv1.GetResponse
		err  error
		mock mocker
	}{
		{
			name: "success user create",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *mockdefserv.MockUserService {
				mocker := mockdefserv.NewMockUserService(t)

				mocker.On("Get", ctx, uint64(req.Id)).Return(userModel, nil)

				return mocker
			},
		},
		{
			name: "fail user create",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			mock: func() *mockdefserv.MockUserService {
				mocker := mockdefserv.NewMockUserService(t)

				mocker.On("Get", ctx, uint64(req.Id)).Return(model.User{}, serviceErr)

				return mocker
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mock()

			impl := user.NewImplementation(slog.New(slogpretty.NewHandler()), mocker)

			output, err := impl.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}
