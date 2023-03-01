package server_crud

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "playlist/proto"
)

func (s *Server) Next(context.Context, *emptypb.Empty) (*pb.NextStatus, error) {
	if s.PlayList.CurrentCursor.Next == nil {
		return &pb.NextStatus{Error: "no next tracks"}, fmt.Errorf("no next tracks")
	} else {
		s.PlayList.CurrentCursor = s.PlayList.CurrentCursor.Next
	}
	return &pb.NextStatus{Error: "you are on next track"}, nil
}
