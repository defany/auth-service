package usertests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/auth-service/app/internal/repository"
	mockrepository "github.com/defany/auth-service/app/internal/repository/mocks"
	userservice "github.com/defany/auth-service/app/internal/service/user"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/auth-service/app/pkg/hasher"
	"github.com/defany/db/pkg/postgres"
	mockpostgres "github.com/defany/db/pkg/postgres/mocks"
	"github.com/defany/slogger/pkg/logger/sl"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_SuccessUserCreate(t *testing.T) {
	type args struct {
		ctx             context.Context
		userCreateInput model.UserCreate
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(false, true, true, false, false, 6)
		role     = userv1.UserRole_name[int32(userv1.UserRole_ADMIN)]

		userCreateInput = model.UserCreate{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}

		logCreateInput = model.Log{
			Action: model.LogCreateUser,
			UserID: userID,
		}
	)

	tests := []struct {
		name   string
		args   args
		want   uint64
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "success",
			args: args{
				ctx:             context.Background(),
				userCreateInput: userCreateInput,
				logCreateInput:  logCreateInput,
			},
			want: userID,
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

				userRepo.On("Create", txCtx, tt.userCreateInput).Return(userID, nil)

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

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log, hasher.NewPasswordMock(tt.args.userCreateInput.Password))

			output, err := service.Create(tt.args.ctx, tt.args.userCreateInput)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}

func TestService_FailUserCreateProcessTx(t *testing.T) {
	type args struct {
		ctx             context.Context
		userCreateInput model.UserCreate
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		err   = errors.New("some error")
		slErr = sl.Err("service.Create", err)
	)

	tests := []struct {
		name   string
		args   args
		want   uint64
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "failed to start tx because ReadCommitted returned an error",
			args: args{
				ctx:             context.Background(),
				userCreateInput: model.UserCreate{},
				logCreateInput:  model.Log{},
			},
			want: uint64(0),
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

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log, hasher.NewPasswordMock(tt.args.userCreateInput.Password))

			output, err := service.Create(tt.args.ctx, tt.args.userCreateInput)

			require.Error(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}

func TestService_FailUserCreate(t *testing.T) {
	type args struct {
		ctx             context.Context
		userCreateInput model.UserCreate
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(false, true, true, false, false, 6)
		role     = userv1.UserRole_name[int32(userv1.UserRole_ADMIN)]

		userCreateInput = model.UserCreate{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}

		logCreateInput = model.Log{
			Action: model.LogCreateUser,
			UserID: userID,
		}

		err   = errors.New("some error")
		slErr = sl.Err("service.Create", err)
	)

	tests := []struct {
		name   string
		args   args
		want   uint64
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "fail user create because user repository returned an error",
			args: args{
				ctx:             context.Background(),
				userCreateInput: userCreateInput,
				logCreateInput:  logCreateInput,
			},
			want: uint64(0),
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
				user.On("Create", txCtx, tt.userCreateInput).Return(uint64(0), err)

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

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log, hasher.NewPasswordMock(tt.args.userCreateInput.Password))

			output, err := service.Create(tt.args.ctx, tt.args.userCreateInput)

			require.Error(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}

func TestService_FailUserCreateLog(t *testing.T) {
	type args struct {
		ctx             context.Context
		userCreateInput model.UserCreate
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(false, true, true, false, false, 6)
		role     = userv1.UserRole_name[int32(userv1.UserRole_ADMIN)]

		userCreateInput = model.UserCreate{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}

		logCreateInput = model.Log{
			Action: model.LogCreateUser,
			UserID: userID,
		}

		err   = errors.New("some error")
		slErr = sl.Err("service.Create", err)
	)

	tests := []struct {
		name   string
		args   args
		want   uint64
		err    error
		mocker func(tt args) mocker
	}{
		{
			name: "fail log create because log repository returned an error",
			args: args{
				ctx:             context.Background(),
				userCreateInput: userCreateInput,
				logCreateInput:  logCreateInput,
			},
			want: uint64(0),
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
				user.On("Create", txCtx, tt.userCreateInput).Return(userID, nil)

				log := mockrepository.NewMockLogRepository(t)
				log.On("Log", txCtx, tt.logCreateInput).Return(err)

				return mocker{
					txManager: txManager,
					user:      user,
					log:       log,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Parallel()

		t.Run(tt.name, func(t *testing.T) {
			mocker := tt.mocker(tt.args)

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log, hasher.NewPasswordMock(tt.args.userCreateInput.Password))

			output, err := service.Create(tt.args.ctx, tt.args.userCreateInput)

			require.Error(t, tt.err, err)
			require.Equal(t, tt.want, output)
		})
	}
}
