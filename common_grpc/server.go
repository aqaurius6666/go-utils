package commongrpc

import (
	"context"
	"os"
	"time"

	"github.com/aqaurius6666/go-utils/common_grpc/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	_ healthpb.HealthServer = (*CommonServer)(nil)
	_ pb.CommonServer       = (*CommonServer)(nil)
)

type CommonServer struct {
	pb.UnimplementedCommonServer
	Logger    *logrus.Logger
	AllowKill bool
}

func NewCommonServer(logger *logrus.Logger, allowKill bool) *CommonServer {
	return &CommonServer{
		Logger:    logger,
		AllowKill: allowKill,
	}
}

func (c *CommonServer) Echo(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (c *CommonServer) Check(_ context.Context, request *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (c *CommonServer) Watch(request *healthpb.HealthCheckRequest, server healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}

func (c *CommonServer) Kill(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	c.Logger.Warnf("Recieved shutting down request at %s", time.Now())
	if c.AllowKill {
		c.Logger.Warnf("Accepted shutting down request at %s", time.Now())
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			return nil, err
		}
		err = process.Kill()
		if err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	}
	c.Logger.Warnf("Rejected shutting down request at %s", time.Now())
	return nil, status.Errorf(codes.PermissionDenied, "no permission")
}
