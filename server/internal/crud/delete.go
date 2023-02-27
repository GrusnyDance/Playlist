package crud

import (
	"context"
	"fmt"
	pb "playlist/proto"
)

func (s *Server) DeleteSong(ctx context.Context, in *pb.SongName) (*pb.DeleteStatus, error) {
	if s.PlayList.CurrentPlay.Name == in.Name {
		fmt.Println("hahaha")
	}
	return &pb.DeleteStatus{Error: "am fine thanks"}, nil
}
