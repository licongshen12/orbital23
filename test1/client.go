package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	Arithmetic "github.com/licongshen12/orbital23/gen-go/arithmetic" // Replace with your actual module path
)

func handleClient(client *Arithmetic.ArithmeticClient) (err error) {
	ctx := context.Background()

	res, err := client.Add(ctx, 10, 20)
	if err != nil {
		fmt.Println("Error calling Add:", err)
		return err
	}

	fmt.Println("Result:", res)
	return nil
}

func main() {
	var protocolFactory thrift.TProtocolFactory
	var transportFactory thrift.TTransportFactory

	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory = thrift.NewTTransportFactory()

	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return
	}
	defer transport.Close()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("Error getting Transport:", err)
		return
	}

	if err := transport.Open(); err != nil {
		fmt.Println("Error opening transport:", err)
		return
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := Arithmetic.NewArithmeticClient(thrift.NewTStandardClient(iprot, oprot))

	if err := handleClient(client); err != nil {
		fmt.Println("Error handling client:", err)
	}
}
