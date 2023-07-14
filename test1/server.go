// server.go
package main

import (
	"context"
	"fmt"
	arithmetic "github.com/licongshen12/orbital23/tree/main/test1/gen-go/arithmetic"
	"github.com/apache/thrift/lib/go/thrift"
)

type arithmeticHandler struct{}

func (p *arithmeticHandler) Add(ctx context.Context, num1 int32, num2 int32) (int32, error) {
	return num1 + num2, nil
}

func RunServer() {
	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		fmt.Println("Error!", err)
		return
	}

	handler := &arithmeticHandler{}
	processor := arithmetic.NewArithmeticProcessor(handler)

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", ":9090")
	server.Serve()
}

func main() {
	RunServer()
}
