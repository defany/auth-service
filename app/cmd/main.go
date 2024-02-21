package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/defany/auth-service/app/config"
	"github.com/defany/auth-service/app/pkg/gen/user/v1"
	"github.com/defany/auth-service/app/pkg/logger/sl"
	"github.com/defany/auth-service/app/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const port = 50001

const (
	Users = "users"
)

const (
	UsersID              = "id"
	UsersName            = "name"
	UsersEmail           = "email"
	UsersPassword        = "password"
	UsersPasswordConfirm = "password_confirm"
	UsersRole            = "role"
	UsersCreatedAt       = "created_at"
	UsersUpdatedAt       = "updated_at"
)

var (
	ErrCreateUser = errors.New("failed to create user")
	ErrGetUser    = errors.New("failed to get user")
	ErrUpdateUser = errors.New("failed to update user")
	ErrDeleteUser = errors.New("failed to delete user")
)

type User struct {
	ID              uint64    `db:"id"`
	Name            string    `db:"name"`
	Email           string    `db:"email"`
	Password        string    `db:"password"`
	PasswordConfirm string    `db:"password_confirm"`
	Role            string    `db:"role"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type server struct {
	userv1.UnimplementedUserServiceServer

	db  *pgxpool.Pool
	qb  squirrel.StatementBuilderType
	log *slog.Logger
}

func (s *server) Create(ctx context.Context, request *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	log := s.log.With(
		slog.String(UsersName, request.GetName()),
		slog.String(UsersEmail, request.GetEmail()),
		slog.String(UsersPassword, request.GetPassword()),
		slog.String(UsersPasswordConfirm, request.GetPasswordConfirm()),
		slog.String(UsersRole, request.GetRole().String()),
	)

	log.Info("create user request")

	q := s.qb.Insert(Users).
		Columns(UsersName, UsersEmail, UsersPassword, UsersPasswordConfirm, UsersRole).
		Values(request.GetName(), request.GetEmail(), request.GetPassword(), request.GetPasswordConfirm(), request.GetRole()).
		Suffix("returning id")

	sql, args, err := q.ToSql()
	if err != nil {
		log.Error("failed build query to create user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrCreateUser.Error())
	}

	rows, err := s.db.Query(ctx, sql, args...)
	if err != nil {
		log.Error("failed to execute query to create user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrCreateUser.Error())
	}

	id, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		log.Error("failed to collect user id from db", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrCreateUser.Error())
	}

	return &userv1.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Get(ctx context.Context, request *userv1.GetRequest) (*userv1.GetResponse, error) {
	log := s.log.With(
		slog.Int64(UsersID, request.GetId()),
	)

	log.Info("get user request")

	q := s.qb.Select(UsersID, UsersEmail, UsersRole, UsersCreatedAt, UsersUpdatedAt).
		From(Users).
		Where(squirrel.Eq{
			UsersID: request.GetId(),
		})

	sql, args, err := q.ToSql()
	if err != nil {
		log.Error("failed build query to get user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrGetUser.Error())
	}

	rows, err := s.db.Query(ctx, sql, args...)
	if err != nil {
		log.Error("failed to execute query to get user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrGetUser.Error())
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		log.Error("failed to collect user from db", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrGetUser.Error())
	}

	roleNum := userv1.UserRole_value[user.Role]

	resData := userv1.GetResponse{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Role:      userv1.UserRole(roleNum),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	log.Info("user info", slog.Any("info", &resData))

	return &resData, nil
}

func (s *server) Update(ctx context.Context, request *userv1.UpdateRequest) (*emptypb.Empty, error) {
	log := s.log.With(
		slog.Int64(UsersID, request.GetId()),
		slog.String(UsersEmail, request.GetEmail().GetValue()),
		slog.String(UsersName, request.GetName().GetValue()),
		slog.String(UsersRole, request.GetRole().String()),
	)

	log.Info("update user request")

	q := s.qb.Update(Users).Where(squirrel.Eq{
		UsersID: request.GetId(),
	})

	if request.GetEmail() != nil {
		q = q.Set(UsersEmail, request.GetEmail().GetValue())
	}

	if request.GetName() != nil {
		q = q.Set(UsersName, request.GetName().GetValue())
	}

	if request.GetRole() != userv1.UserRole_UNKNOWN {
		q = q.Set(UsersRole, request.GetRole().String())
	}

	sql, args, err := q.ToSql()
	if err != nil {
		log.Error("failed build query to update user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrUpdateUser.Error())
	}

	_, err = s.db.Exec(ctx, sql, args...)
	if err != nil {
		log.Error("failed to execute query to update user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrUpdateUser.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, request *userv1.DeleteRequest) (*emptypb.Empty, error) {
	log := s.log.With(
		slog.Int64("id", request.GetId()),
	)

	log.Info("delete user request")

	q := s.qb.Delete(Users).Where(squirrel.Eq{
		UsersID: request.GetId(),
	})

	sql, args, err := q.ToSql()
	if err != nil {
		log.Error("failed build query to delete user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrDeleteUser.Error())
	}

	_, err = s.db.Exec(ctx, sql, args...)
	if err != nil {
		log.Error("failed to execute query to delete user", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, ErrDeleteUser.Error())
	}

	return &emptypb.Empty{}, nil
}

func main() {
	cfg := config.MustLoad()

	log := sl.NewSlogLogger(cfg.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := postgres.NewConfig(cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)

	db, err := postgres.NewClient(ctx, log, cfg.Database.ConnectAttempts, cfg.Database.ConnectAttemptsDelay, dbConfig)
	if err != nil {
		log.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error("failed to listen", slog.String("error", err.Error()))
		os.Exit(1)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	userv1.RegisterUserServiceServer(s, &server{
		db:  db,
		qb:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		log: log,
	})

	log.Info("listening", slog.String("port", lis.Addr().String()))

	if err := s.Serve(lis); err != nil {
		log.Error("failed to server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
