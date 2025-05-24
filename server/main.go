package main

import (
 "context"
 pb "github.com/Nipun2001M/Simple-gRPC-Server/helloworld"
 "google.golang.org/grpc"
 "log"
 "net"
)

type server struct{
	pb.UnimplementedHelloworldServiceServer
}

func (s * server) SayHello(ctx context.Context,in *pb.HelloWorldRequest) (*pb.HelloWorldResponse,error){
	return &pb.HelloWorldResponse{
		Message: "Hello World",
	},nil
}

func main(){
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil{
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s:=grpc.NewServer()
	pb.RegisterHelloworldServiceServer(s,&server{})
	log.Printf("gRPC server listning at %v",lis.Addr())
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to serve %v", err)
	}

}

