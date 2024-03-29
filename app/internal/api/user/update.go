package user

import (
	"context"
	"log/slog"

	"github.com/defany/auth-service/app/internal/converter"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/slogger/pkg/logger/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *userv1.UpdateRequest) (*emptypb.Empty, error) {
	log := i.log.With(slog.String("op", sl.FnName()))

	err := i.service.Update(ctx, converter.ToUserUpdate(req))
	if err != nil {
		log.Error("failed to update user", sl.ErrAttr(err))

		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return &emptypb.Empty{}, nil
}
