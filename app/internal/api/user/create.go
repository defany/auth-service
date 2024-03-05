package user

import (
	"context"
	"github.com/defany/auth-service/app/internal/converter"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/auth-service/app/pkg/logger/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (i *Implementation) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	log := i.log.With(
		slog.String("op", sl.FnName()),
	)

	userID, err := i.service.Create(ctx, converter.ToUserCreate(req))
	if err != nil {
		log.Error("failed to create user", sl.ErrAttr(err))

		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userv1.CreateResponse{
		Id: int64(userID),
	}, nil
}
