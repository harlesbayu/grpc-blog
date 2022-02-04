package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/harlesbayu/grpc-blog/blogpb"
	"google.golang.org/grpc"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog Service Started")

	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the server")
	lis.Close()
	fmt.Println("End of program")
}
