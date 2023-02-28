package crud

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "playlist/proto"
)

func (s *Server) Prev(context.Context, *emptypb.Empty) (*pb.PrevStatus, error) {
	if s.PlayList.CurrentCursor.Prev == nil {
		return &pb.PrevStatus{Error: "no previous tracks"}, fmt.Errorf("no previous tracks")
	} else {
		s.PlayList.CurrentCursor = s.PlayList.CurrentCursor.Prev
	}
	return &pb.PrevStatus{Error: "you are on previous track"}, nil
}
