package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"os"
	"os/signal"
	"playlist/client/usecase"
	pb "playlist/proto"
	"syscall"
)

func main() {
	conn, err := grpc.Dial("localhost:9999",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlaylistClient(conn)
	usecase.New(&client)

	//// usecase 2
	//empty := &emptypb.Empty{}
	////_, err = client.Play(context.Background(), empty, grpc.StreamInterceptor(NewClientStreamInterceptor()))
	//stream, err := client.Play(context.Background(), empty)
	//if err != nil {
	//	fmt.Println("error while playing", err)
	//}
	//go internal.PlaySound(&stream)
	//
	//// usecase 2

	_, cancel := context.WithCancel(context.Background())
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		// сначала вызвать паузу
		cancel()
	}
}
