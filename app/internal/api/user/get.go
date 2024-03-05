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

func (i *Implementation) Get(ctx context.Context, request *userv1.GetRequest) (*userv1.GetResponse, error) {
	op := sl.FnName()

	log := i.log.With(slog.String("op", op))

	log.Info("getting user by id")

	output, err := i.service.Get(ctx, uint64(request.GetId()))
	if err != nil {
		log.Error("failed to get user by id", sl.OpErrAttr(op, err))

		// Дождусь твоего варианта обработки ошибок от бизнес логики, мой мне не нравится, поэтому пока буду отдавать всегда Internal
		return nil, status.Error(codes.Internal, "failed to get user by id")
	}

	return converter.ToGetResponse(output), nil
}
