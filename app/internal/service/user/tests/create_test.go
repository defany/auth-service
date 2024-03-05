package usertests

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/model"
	mockrepository "github.com/defany/auth-service/app/internal/repository/mocks"
	userservice "github.com/defany/auth-service/app/internal/service/user"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	mockpostgres "github.com/defany/auth-service/app/pkg/postgres/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_Create(t *testing.T) {
	type userMocker func() *mockrepository.MockUserRepository
	type logMocker func() *mockrepository.MockLogRepository
	type txMocker func() *mockpostgres.MockTxManager

	type args struct {
		ctx context.Context
		req model.UserCreate
	}

	var (
		ctx = context.Background()

		//id = gofakeit.Uint64()

		serviceErr = fmt.Errorf("failed to create user")

		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(false, true, true, true, false, 6)
		hashedPassword  = md5.Sum([]byte(password))
		passwordConfirm = hex.EncodeToString(hashedPassword[:])
		role            = userv1.UserRole(int32(gofakeit.Float64Range(0, 2)))

		input = model.UserCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role.String(),
		}

		//logInput = model.Log{
		//	Action: model.LogCreateUser,
		//	UserID: id,
		//}
	)

	tests := []struct {
		name     string
		args     args
		want     uint64
		err      error
		userMock userMocker
		logMock  logMocker
		txMock   txMocker
	}{
		{
			name: "read committed exec success",
			args: args{
				ctx: ctx,
				req: input,
			},
			want: 0,
			err:  nil,
			userMock: func() *mockrepository.MockUserRepository {
				mocker := mockrepository.NewMockUserRepository(t)

				return mocker
			},
			logMock: func() *mockrepository.MockLogRepository {
				mocker := mockrepository.NewMockLogRepository(t)

				return mocker
			},
			txMock: func() *mockpostgres.MockTxManager {
				mocker := mockpostgres.NewMockTxManager(t)

				mocker.On("ReadCommitted", ctx, mock.AnythingOfType("postgres.Handler")).Return(nil)

				return mocker
			},
		},
		{
			name: "repo mock",
			args: args{
				ctx: ctx,
				req: input,
			},
			want: 0,
			err:  serviceErr,
			userMock: func() *mockrepository.MockUserRepository {
				mocker := mockrepository.NewMockUserRepository(t)

				mocker.On("Create", ctx, input).Return(0, serviceErr)

				return mocker
			},
			logMock: func() *mockrepository.MockLogRepository {
				mocker := mockrepository.NewMockLogRepository(t)

				return mocker
			},
			txMock: func() *mockpostgres.MockTxManager {
				mocker := mockpostgres.NewMockTxManager(t)

				mocker.On("ReadCommitted", ctx, mock.AnythingOfType("postgres.Handler")).Return(serviceErr)

				return mocker
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userMock := tt.userMock()
			logMock := tt.logMock()
			txMock := tt.txMock()

			service := userservice.NewService(txMock, userMock, logMock)

			userID, err := service.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, userID)
		})
	}
}
