package user

import (
	"context"
	"log/slog"

	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/defany/slogger/pkg/logger/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *userv1.DeleteRequest) (*emptypb.Empty, error) {
	log := i.log.With(slog.String("op", sl.FnName()))

	err := i.service.Delete(ctx, uint64(req.GetId()))
	if err != nil {
		log.Error("failed to delete user", sl.ErrAttr(err))

		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return &emptypb.Empty{}, nil
}
