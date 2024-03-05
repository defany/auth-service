package usertests

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/api/user"
	mockdefserv "github.com/defany/auth-service/app/internal/service/mocks"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/auth-service/app/pkg/logger/handlers/slogpretty"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
	"testing"
)

func TestImplementation_Delete(t *testing.T) {
	type mocker func() *mockdefserv.MockUserService

	type args struct {
		ctx context.Context
		req *userv1.DeleteRequest
	}

	var (
		ctx = context.Background()

		serviceErr = status.Error(codes.Internal, "failed to delete user")

		req = &userv1.DeleteRequest{
			Id: gofakeit.Int64(),
		}

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
			name: "success user create",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *mockdefserv.MockUserService {
				mocker := mockdefserv.NewMockUserService(t)

				mocker.On("Delete", ctx, uint64(req.Id)).Return(nil)

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

				mocker.On("Delete", ctx, uint64(req.Id)).Return(serviceErr)

				return mocker
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mock()

			impl := user.NewImplementation(slog.New(slogpretty.NewHandler()), mocker)

			output, err := impl.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}
