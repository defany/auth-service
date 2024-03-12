package usertests

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/auth-service/app/internal/repository"
	mockrepository "github.com/defany/auth-service/app/internal/repository/mocks"
	userservice "github.com/defany/auth-service/app/internal/service/user"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/auth-service/app/pkg/hash"
	"github.com/defany/auth-service/app/pkg/logger/sl"
	"github.com/defany/auth-service/app/pkg/postgres"
	mockpostgres "github.com/defany/auth-service/app/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_SuccessUserGet(t *testing.T) {
	type args struct {
		ctx            context.Context
		userGetInput   uint64
		userGetOutput  model.User
		logCreateInput model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(false, true, true, false, false, 6)
		passwordConfirm = hash.MD5(password)
		role            = userv1.UserRole_name[int32(userv1.UserRole_ADMIN)]
		createdAt       = gofakeit.Date()
		updatedAt       = gofakeit.Date()

		userGetInput = userID

		userGetOutput = model.User{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}

		logCreateInput = model.Log{
			Action: model.LogGetUser,
			UserID: userID,
		}
	)

	tests := []struct {
		name   string
		args   args
		want   model.User
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "success",
			args: args{
				ctx:            context.Background(),
				userGetInput:   userGetInput,
				userGetOutput:  userGetOutput,
				logCreateInput: logCreateInput,
			},
			want: userGetOutput,
			err:  nil,
			mocker: func(tt args) mocker {
				txOpts := pgx.TxOptions{
					IsoLevel: pgx.ReadCommitted,
				}

				tx := mockpostgres.NewMockTx(t)

				txCtx := postgres.InjectTX(tt.ctx, tx)

				tx.On("Commit", txCtx).Return(nil)

				db := mockpostgres.NewMockPostgres(t)
				db.On("BeginTx", tt.ctx, txOpts).Return(tx, nil)

				txManager := postgres.NewTxManager(db)
				userRepo := mockrepository.NewMockUserRepository(t)
				logRepo := mockrepository.NewMockLogRepository(t)

				userRepo.On("User", txCtx, tt.userGetInput).Return(tt.userGetOutput, nil)

				logRepo.On("Log", txCtx, tt.logCreateInput).Return(nil)

				return mocker{
					txManager: txManager,
					user:      userRepo,
					log:       logRepo,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Parallel()

		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mocker(tt.args)

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log)

			output, err := service.Get(tt.args.ctx, tt.args.userGetInput)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}

func TestService_FailUserGetProcessTx(t *testing.T) {
	type args struct {
		ctx            context.Context
		userGetInput   uint64
		userGetOutput  model.User
		logCreateInput model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		err   = errors.New("some error")
		slErr = sl.Err("service.Get", err)

		userGetOutput = model.User{}
	)

	tests := []struct {
		name   string
		args   args
		want   model.User
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "failed to start tx because ReadCommitted returned an error",
			args: args{
				ctx:            context.Background(),
				userGetInput:   0,
				userGetOutput:  userGetOutput,
				logCreateInput: model.Log{},
			},
			want: userGetOutput,
			err:  slErr,
			mocker: func(tt args) mocker {
				txManager := mockpostgres.NewMockTxManager(t)
				txManager.On("ReadCommitted", tt.ctx, mock.AnythingOfType("postgres.Handler")).Return(err)

				return mocker{
					txManager: txManager,
					user:      nil,
					log:       nil,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Parallel()

		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mocker(tt.args)

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log)

			output, err := service.Get(tt.args.ctx, tt.args.userGetInput)

			require.Error(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}

func TestService_FailUserGet(t *testing.T) {
	type args struct {
		ctx            context.Context
		userGetInput   uint64
		userGetOutput  model.User
		logCreateInput model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		userGetInput = userID

		userGetOutput = model.User{}

		logCreateInput = model.Log{
			Action: model.LogGetUser,
			UserID: userID,
		}

		err   = errors.New("some error")
		slErr = sl.Err("service.Get", err)
	)

	tests := []struct {
		name   string
		args   args
		want   model.User
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "fail user create because user repository returned an error",
			args: args{
				ctx:            context.Background(),
				userGetInput:   userGetInput,
				userGetOutput:  userGetOutput,
				logCreateInput: logCreateInput,
			},
			want: model.User{},
			err:  slErr,
			mocker: func(tt args) mocker {
				txOpts := pgx.TxOptions{
					IsoLevel: pgx.ReadCommitted,
				}

				tx := mockpostgres.NewMockTx(t)

				txCtx := postgres.InjectTX(tt.ctx, tx)

				tx.On("Rollback", txCtx).Return(nil)

				db := mockpostgres.NewMockPostgres(t)
				db.On("BeginTx", tt.ctx, txOpts).Return(tx, nil)

				txManager := postgres.NewTxManager(db)

				user := mockrepository.NewMockUserRepository(t)
				user.On("User", txCtx, tt.userGetInput).Return(userGetOutput, err)

				return mocker{
					txManager: txManager,
					user:      user,
					log:       nil,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Parallel()

		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mocker(tt.args)

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log)

			output, err := service.Get(tt.args.ctx, tt.args.userGetInput)

			require.Error(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}

func TestService_FailUserGetLog(t *testing.T) {
	type args struct {
		ctx            context.Context
		userGetInput   uint64
		userGetOutput  model.User
		logCreateInput model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		userGetInput = userID

		userGetOutput = model.User{}

		logCreateInput = model.Log{
			Action: model.LogGetUser,
			UserID: userID,
		}

		err   = errors.New("some error")
		slErr = sl.Err("service.Get", err)
	)

	tests := []struct {
		name   string
		args   args
		want   model.User
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "success",
			args: args{
				ctx:            context.Background(),
				userGetInput:   userGetInput,
				userGetOutput:  userGetOutput,
				logCreateInput: logCreateInput,
			},
			want: userGetOutput,
			err:  slErr,
			mocker: func(tt args) mocker {
				txOpts := pgx.TxOptions{
					IsoLevel: pgx.ReadCommitted,
				}

				tx := mockpostgres.NewMockTx(t)

				txCtx := postgres.InjectTX(tt.ctx, tx)

				tx.On("Rollback", txCtx).Return(nil)

				db := mockpostgres.NewMockPostgres(t)
				db.On("BeginTx", tt.ctx, txOpts).Return(tx, nil)

				txManager := postgres.NewTxManager(db)
				userRepo := mockrepository.NewMockUserRepository(t)
				logRepo := mockrepository.NewMockLogRepository(t)

				userRepo.On("User", txCtx, tt.userGetInput).Return(tt.userGetOutput, nil)

				logRepo.On("Log", txCtx, tt.logCreateInput).Return(err)

				return mocker{
					txManager: txManager,
					user:      userRepo,
					log:       logRepo,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Parallel()

		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mocker(tt.args)

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log)

			output, err := service.Get(tt.args.ctx, tt.args.userGetInput)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}
