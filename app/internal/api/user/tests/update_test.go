package usertests

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/api/user"
	"github.com/defany/auth-service/app/internal/converter"
	mockdefserv "github.com/defany/auth-service/app/internal/service/mocks"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/slogger/pkg/logger/handlers/slogpretty"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
	"testing"
)

func TestImplementation_Update(t *testing.T) {
	type mocker func() *mockdefserv.MockUserService

	type args struct {
		ctx context.Context
		req *userv1.UpdateRequest
	}

	var (
		ctx = context.Background()

		id = gofakeit.Uint64()

		serviceErr = status.Error(codes.Internal, "failed to update user")

		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = userv1.UserRole(int32(gofakeit.Float64Range(0, 2)))

		req = &userv1.UpdateRequest{
			Id:    int64(id),
			Name:  &name,
			Email: &email,
			Role:  &role,
		}

		input = converter.ToUserUpdate(req)

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name string
		args args
		want *emptypb.Empty
		err  error
		mock mocker
	}{
		{
			name: "user updated successfully",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *mockdefserv.MockUserService {
				mocker := mockdefserv.NewMockUserService(t)

				mocker.On("Update", ctx, input).Return(nil)

				return mocker
			},
		},
		{
			name: "failed to update user",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			mock: func() *mockdefserv.MockUserService {
				mocker := mockdefserv.NewMockUserService(t)

				mocker.On("Update", ctx, input).Return(serviceErr)

				return mocker
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mock()

			impl := user.NewImplementation(slog.New(slogpretty.NewHandler()), mocker)

			output, err := impl.Update(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}
