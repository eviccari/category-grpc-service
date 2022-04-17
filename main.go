package main

import (
	"fmt"

	"google.golang.org/grpc"
)

func main()  {
	grpcServer := grpc.NewServer()
	fmt.Println(grpcServer)
}