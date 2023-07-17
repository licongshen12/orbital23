package main

import (
	"context"
	calculator "rpctest2/kitex_gen/calculator"
)

// CalculatorImpl implements the last service interface defined in the IDL.
type CalculatorImpl struct{}

// Add implements the CalculatorImpl interface.
func (s *CalculatorImpl) Add(ctx context.Context, num1 *calculator.Request, num2 *calculator.Request) (resp *calculator.Response, err error) {
	// TODO: Your code here...
	return &calculator.Response{Result_: num1.Num1 + num1.Num2}, nil
}
