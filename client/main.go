package main

import (
	"context"
	pb "github.com/Nipun2001M/Simple-gRPC-Server/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}

	defer conn.Close()
	c := pb.NewHelloworldServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r,err:=c.SayHello(ctx,&pb.HelloWorldRequest{})
	if err!=nil{
		log.Fatalf("error calling function sayhello %v",err)
	}
	log.Printf("Response from gRPC sayhello function %s",r.GetMessage())

}
