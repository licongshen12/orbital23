package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	Arithmetic "github.com/licongshen12/orbital23/gen-go/arithmetic" // Replace with your actual module path
)

// Make sure the struct implements the Arithmetic interface
type ArithmeticHandler struct{}

func NewArithmeticHandler() *ArithmeticHandler {
	return &ArithmeticHandler{}
}

func (*ArithmeticHandler) Add(ctx context.Context, num1 int32, num2 int32) (int32, error) {
	// Implement the Add operation here
	return num1 + num2, nil
}

func main() {
	handler := NewArithmeticHandler()
	processor := tutorial.NewArithmeticProcessor(handler)

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()

	address := "localhost:9090"
	transport, err := thrift.NewTServerSocket(address)

	if err != nil {
		fmt.Println("Error creating Thrift server socket:", err)
		return
	}

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", address)
	if err := server.Serve(); err != nil {
		fmt.Println("Error starting Thrift server:", err)
	}
}

