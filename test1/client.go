// client.go
package main

import (
	"context"
	"fmt"
	arithmetic "github.com/licongshen12/orbital23/tree/main/test1/gen-go/arithmetic"
	"github.com/apache/thrift/lib/go/thrift"
)

func RunClient() {
	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return
	}
	defer transport.Close()

	transportFactory := thrift.NewTTransportFactory()
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := transport.Open(); err != nil {
		fmt.Println(err)
		return
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := arithmetic.NewArithmeticClientFactory(transport, protocolFactory)

	sum, err := client.Add(context.Background(), 10, 20)
	if err != nil {
		fmt.Println("Error calling Add:", err)
		return
	}

	fmt.Println("Sum:", sum)
}

func main() {
	RunClient()
}
