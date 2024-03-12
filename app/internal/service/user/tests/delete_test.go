package usertests

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit"
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/auth-service/app/internal/repository"
	mockrepository "github.com/defany/auth-service/app/internal/repository/mocks"
	userservice "github.com/defany/auth-service/app/internal/service/user"
	"github.com/defany/auth-service/app/pkg/logger/sl"
	"github.com/defany/auth-service/app/pkg/postgres"
	mockpostgres "github.com/defany/auth-service/app/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_SuccessUserDelete(t *testing.T) {
	type args struct {
		ctx             context.Context
		userDeleteInput uint64
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		logCreateInput = model.Log{
			Action: model.LogDeleteUser,
			UserID: userID,
		}
	)

	tests := []struct {
		name   string
		args   args
		want   error
		mocker func(tt args) mocker
	}{
		{
			name: "success",
			args: args{
				ctx:             context.Background(),
				userDeleteInput: userID,
				logCreateInput:  logCreateInput,
			},
			want: nil,
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

				userRepo.On("Delete", txCtx, tt.userDeleteInput).Return(nil)

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

			err := service.Delete(tt.args.ctx, tt.args.userDeleteInput)

			require.Equal(t, tt.want, err)
		})
	}
}

func TestService_FailUserDeleteProcessTx(t *testing.T) {
	type args struct {
		ctx             context.Context
		userDeleteInput uint64
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		err   = errors.New("some error")
		slErr = sl.Err("service.Delete", err)
	)

	tests := []struct {
		name   string
		args   args
		want   error
		mocker func(tt args) mocker
	}{
		{
			name: "failed to start tx because ReadCommitted returned an error",
			args: args{
				ctx:             context.Background(),
				userDeleteInput: 0,
				logCreateInput:  model.Log{},
			},
			want: slErr,
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

			err := service.Delete(tt.args.ctx, tt.args.userDeleteInput)

			require.Equal(t, tt.want, err)
		})
	}
}

func TestService_FailUserDelete(t *testing.T) {
	type args struct {
		ctx             context.Context
		userDeleteInput uint64
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		err   = errors.New("some error")
		slErr = sl.Err("service.Delete", err)

		logCreateInput = model.Log{
			Action: model.LogCreateUser,
			UserID: userID,
		}
	)

	tests := []struct {
		name   string
		args   args
		want   error
		mocker func(tt args) mocker
	}{
		{
			name: "fail user create because user repository returned an error",
			args: args{
				ctx:             context.Background(),
				userDeleteInput: userID,
				logCreateInput:  logCreateInput,
			},
			want: slErr,
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
				user.On("Delete", txCtx, tt.userDeleteInput).Return(err)

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

			err := service.Delete(tt.args.ctx, tt.args.userDeleteInput)

			require.Equal(t, tt.want, err)
		})
	}
}

func TestService_FailUserDeleteLog(t *testing.T) {
	type args struct {
		ctx             context.Context
		userDeleteInput uint64
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		err   = errors.New("some error")
		slErr = sl.Err("service.Delete", err)

		logCreateInput = model.Log{
			Action: model.LogDeleteUser,
			UserID: userID,
		}
	)

	tests := []struct {
		name   string
		args   args
		want   error
		mocker func(tt args) mocker
	}{
		{
			name: "fail log create because log repository returned an error",
			args: args{
				ctx:             context.Background(),
				userDeleteInput: userID,
				logCreateInput:  logCreateInput,
			},
			want: slErr,
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
				user.On("Delete", txCtx, tt.userDeleteInput).Return(nil)

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

			service := userservice.NewService(mocker.txManager, mocker.user, mocker.log)

			err := service.Delete(tt.args.ctx, tt.args.userDeleteInput)

			require.Equal(t, tt.want, err)
		})
	}
}
