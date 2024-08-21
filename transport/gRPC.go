package transport

import (
	"context"
	"errors"
	"github.com/go-kit/kit/transport/grpc"
	requestdto "gokit/dto/request"
	responsedto "gokit/dto/response"
	endpoints "gokit/endpoint"
	"gokit/proto/pb"
)

type gRPCServer struct {
	add      grpc.Handler
	subtract grpc.Handler
	multiply grpc.Handler
	divide   grpc.Handler
	pb.UnimplementedMathServiceServer
}

func NewGRPCServer(e endpoints.MathEndpoint) pb.MathServiceServer {
	return &gRPCServer{
		add: grpc.NewServer(
			e.AddEndpoint(),
			decodeMathRequest,
			encodeMathResponse,
		),
		subtract: grpc.NewServer(
			e.SubtractEndpoint(),
			decodeMathRequest,
			encodeMathResponse,
		),
		multiply: grpc.NewServer(
			e.MultiplyEndpoint(),
			decodeMathRequest,
			encodeMathResponse,
		),
		divide: grpc.NewServer(
			e.DivideEndpoint(),
			decodeMathRequest,
			encodeMathResponse,
		),
	}
}

func (s gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s gRPCServer) Subtract(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.subtract.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s gRPCServer) Multiply(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.multiply.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s gRPCServer) Divide(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.divide.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*pb.MathRequest)
	if !ok {
		return nil, errors.New("invalid request body")
	}

	return requestdto.MathRequest{
		NumA: req.NumA,
		NumB: req.NumB,
	}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(responsedto.MathResponse)
	if !ok {
		return nil, errors.New("invalid response body")
	}

	return &pb.MathResponse{Result: resp.Result}, nil
}
