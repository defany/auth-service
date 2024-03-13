package usertests

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/api/user"
	"github.com/defany/auth-service/app/internal/converter"
	mockdefserv "github.com/defany/auth-service/app/internal/service/mocks"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/slogger/pkg/logger/handlers/slogpretty"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"testing"
)

func TestImplementation_Create(t *testing.T) {
	type mocker func() *mockdefserv.MockUserService

	type args struct {
		ctx context.Context
		req *userv1.CreateRequest
	}

	var (
		ctx = context.Background()

		id = gofakeit.Uint64()

		serviceErr = status.Error(codes.Internal, "failed to create user")

		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(false, true, true, true, false, 6)
		hashedPassword  = md5.Sum([]byte(password))
		passwordConfirm = hex.EncodeToString(hashedPassword[:])
		role            = userv1.UserRole(int32(gofakeit.Float64Range(0, 2)))

		req = &userv1.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
		}

		input = converter.ToUserCreate(req)

		res = &userv1.CreateResponse{
			Id: int64(id),
		}
	)

	tests := []struct {
		name string
		args args
		want *userv1.CreateResponse
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

				mocker.On("Create", ctx, input).Return(id, nil)

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

				mocker.On("Create", ctx, input).Return(uint64(0), serviceErr)

				return mocker
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mock()

			impl := user.NewImplementation(slog.New(slogpretty.NewHandler()), mocker)

			output, err := impl.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}
