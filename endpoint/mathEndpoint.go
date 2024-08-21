package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	requestdto "gokit/dto/request"
	responsedto "gokit/dto/response"
	"gokit/service"
)

type MathEndpoint interface {
	AddEndpoint() endpoint.Endpoint
	SubtractEndpoint() endpoint.Endpoint
	MultiplyEndpoint() endpoint.Endpoint
	DivideEndpoint() endpoint.Endpoint
}

type mathEndpoint struct {
	mathService service.MathService
}

func NewMathEndpoint(mathService service.MathService) MathEndpoint {
	return mathEndpoint{
		mathService: mathService,
	}
}

func (e mathEndpoint) AddEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestdto.MathRequest)
		result, err := e.mathService.Add(ctx, req.NumA, req.NumB)
		if err != nil {
			return nil, err
		}
		return responsedto.MathResponse{Result: result}, nil
	}
}

func (e mathEndpoint) SubtractEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestdto.MathRequest)
		result, err := e.mathService.Subtract(ctx, req.NumA, req.NumB)
		if err != nil {
			return nil, err
		}
		return responsedto.MathResponse{Result: result}, nil
	}
}

func (e mathEndpoint) MultiplyEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestdto.MathRequest)
		result, err := e.mathService.Multiply(ctx, req.NumA, req.NumB)
		if err != nil {
			return nil, err
		}
		return responsedto.MathResponse{Result: result}, nil
	}
}

func (e mathEndpoint) DivideEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestdto.MathRequest)
		result, err := e.mathService.Divide(ctx, req.NumA, req.NumB)
		if err != nil {
			return nil, err
		}
		return responsedto.MathResponse{Result: result}, nil
	}
}
