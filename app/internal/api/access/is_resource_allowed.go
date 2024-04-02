package access

import (
	"context"
	"errors"
	"strings"

	accessv1 "github.com/defany/auth-service/app/pkg/gen/proto/access/v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	// TODO: вынести в другое место, если понадобится
	authPrefix = "Bearer "
)

func (i *Implementation) Check(ctx context.Context, req *accessv1.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authInfo := md.Get("authorization")
	if len(authInfo) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	authHeader := authInfo[0]

	if !strings.HasPrefix(authHeader, authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader, authPrefix)

	err := i.service.DoesHaveAccess(ctx, accessToken, req.GetEndpoint())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
