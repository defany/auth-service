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
	"github.com/defany/auth-service/app/pkg/logger/sl"
	"github.com/defany/auth-service/app/pkg/postgres"
	mockpostgres "github.com/defany/auth-service/app/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_SuccessUserUpdate(t *testing.T) {
	type args struct {
		ctx             context.Context
		userUpdateInput model.UserUpdate
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = userv1.UserRole_name[int32(userv1.UserRole_ADMIN)]

		userUpdateInput = model.UserUpdate{
			ID:    userID,
			Name:  &name,
			Email: &email,
			Role:  &role,
		}

		logCreateInput = model.Log{
			Action: model.LogUpdateUser,
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
				userUpdateInput: userUpdateInput,
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

				userRepo.On("Update", txCtx, tt.userUpdateInput).Return(nil)

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

			err := service.Update(tt.args.ctx, tt.args.userUpdateInput)

			require.Equal(t, tt.want, err)
		})
	}
}

func TestService_FailUserUpdateProcessTx(t *testing.T) {
	type args struct {
		ctx             context.Context
		userUpdateInput model.UserUpdate
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		err   = errors.New("some error")
		slErr = sl.Err("service.Update", err)

		userUpdateInput = model.UserUpdate{}
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
				userUpdateInput: userUpdateInput,
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

			err := service.Update(tt.args.ctx, tt.args.userUpdateInput)

			require.Equal(t, tt.want, err)
		})
	}
}

func TestService_FailUserUpdate(t *testing.T) {
	type args struct {
		ctx             context.Context
		userUpdateInput model.UserUpdate
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = userv1.UserRole_name[int32(userv1.UserRole_ADMIN)]

		userUpdateInput = model.UserUpdate{
			ID:    userID,
			Name:  &name,
			Email: &email,
			Role:  &role,
		}

		logCreateInput = model.Log{
			Action: model.LogUpdateUser,
			UserID: userID,
		}

		err   = errors.New("some error")
		slErr = sl.Err("service.Update", err)
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
				userUpdateInput: userUpdateInput,
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
				userRepo := mockrepository.NewMockUserRepository(t)
				logRepo := mockrepository.NewMockLogRepository(t)

				userRepo.On("Update", txCtx, tt.userUpdateInput).Return(err)

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

			err := service.Update(tt.args.ctx, tt.args.userUpdateInput)

			require.Equal(t, tt.want, err)
		})
	}
}

func TestService_FailUserUpdateLog(t *testing.T) {
	type args struct {
		ctx             context.Context
		userUpdateInput model.UserUpdate
		logCreateInput  model.Log
	}

	type mocker struct {
		txManager postgres.TxManager
		user      repository.UserRepository
		log       repository.LogRepository
	}

	var (
		userID = gofakeit.Uint64()

		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = userv1.UserRole_name[int32(userv1.UserRole_ADMIN)]

		userUpdateInput = model.UserUpdate{
			ID:    userID,
			Name:  &name,
			Email: &email,
			Role:  &role,
		}

		logCreateInput = model.Log{
			Action: model.LogUpdateUser,
			UserID: userID,
		}

		err   = errors.New("some error")
		slErr = sl.Err("service.Update", err)
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
				userUpdateInput: userUpdateInput,
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
				userRepo := mockrepository.NewMockUserRepository(t)
				logRepo := mockrepository.NewMockLogRepository(t)

				userRepo.On("Update", txCtx, tt.userUpdateInput).Return(nil)

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

			err := service.Update(tt.args.ctx, tt.args.userUpdateInput)

			require.Equal(t, tt.want, err)
		})
	}
}
