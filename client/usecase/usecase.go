package usecase

import (
	"context"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"playlist/client/internal"
	pb "playlist/proto"
)

func New(client *pb.PlaylistClient) {
	//for i := 0; i < 5; i++ {
	//	go func(j int) {
	//		sn := &pb.SongName{
	//			Name: "валерий меладзе в первый день весны",
	//		}
	//		response, err := (*client).AddSong(context.Background(), sn)
	//		if err != nil {
	//			grpclog.Errorf("fail to add: %v", response)
	//		}
	//	}(i)
	//}

	//sn := &pb.SongName{
	//	Name: "sinatra",
	//}
	//response, err := (*client).AddSong(context.Background(), sn)
	//if err != nil {
	//	grpclog.Errorf("fail to add: %v", response)
	//}

	empty := &emptypb.Empty{}
	stream, err := (*client).Play(context.Background(), empty)
	if err != nil {
		grpclog.Errorf("error while playing %v", err)
	} else {
		go internal.PlaySound(&stream)
	}
}
