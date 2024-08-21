package service

import (
	"context"
	"errors"
)

type MathService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
	Subtract(ctx context.Context, numA, numB float32) (float32, error)
	Multiply(ctx context.Context, numA, numB float32) (float32, error)
	Divide(ctx context.Context, numA, numB float32) (float32, error)
}

type mathService struct{}

func NewMathService() MathService {
	return &mathService{}
}

func (s mathService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}

func (s mathService) Subtract(ctx context.Context, numA, numB float32) (float32, error) {
	return numA - numB, nil
}

func (s mathService) Multiply(ctx context.Context, numA, numB float32) (float32, error) {
	return numA * numB, nil
}

func (s mathService) Divide(ctx context.Context, numA, numB float32) (float32, error) {
	if numB == 0 {
		return 0, errors.New("divide by zero")
	}
	return numA / numB, nil
}
