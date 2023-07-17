package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"

	// "github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/loadbalance/random"
)

func main() {
	// Create a list of backend addresses
	backendAddresses := []string{
		"backend1:8000",
		"backend2:8000",
		"backend3:8000",
	}

	// Create a load balancer with the random load balancing strategy
	loadBalancer := random.NewRandomLB()

	// Create the RPC client with load balancing enabled
	cli, err := genericclient.NewClient("service-name",
		client.WithLoadBalance(loadBalancer),
		client.WithHostPorts(backendAddresses...),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Make RPC calls using the client
	ctx := context.TODO()
	request := []byte("request payload")
	response, err := cli.GenericCall(ctx, "method-name", request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:", response)
}
