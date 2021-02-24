package server

import (
	"context"
	egrpcpb "github.com/syanhaiD/envoy-grpc-prac/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EGRPCService struct{}

func (es *EGRPCService) Test(_ context.Context, _ *emptypb.Empty) (resp *egrpcpb.TestResponse, err error) {
	resp = &egrpcpb.TestResponse{
		Success: true,
		Mes:     "string!",
	}

	return
}
