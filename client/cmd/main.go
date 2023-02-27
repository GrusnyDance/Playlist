package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	pb "playlist/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:9999",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlaylistClient(conn)

	//// usecase 1
	//sn := &pb.SongName{
	//	Name: "перезаряжай",
	//}
	//response, err := client.AddSong(context.Background(), sn)
	//if err != nil {
	//	grpclog.Fatalf("fail to dial: %v", err)
	//}
	//// usecase 1

	// usecase 2
	sn := &pb.SongName{
		Name: "перезаряжай",
	}
	response, err := client.AddSong(context.Background(), sn)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	// usecase 2

	fmt.Println(response.Error, "error is")
}
