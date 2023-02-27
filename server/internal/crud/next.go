package crud

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "playlist/proto"
)

func (s *Server) Next(context.Context, *emptypb.Empty) (*pb.NextStatus, error) {
	return &pb.NextStatus{Error: "am fine thanks"}, nil
}
