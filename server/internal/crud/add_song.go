package crud

import (
	"context"
	pb "playlist/proto"
)

func (s *internal.Server) AddSong(ctx context.Context, in *pb.SongName) (*pb.AddStatus, error) {
	return &pb.AddStatus{Error: "am fine thanks"}, nil
}
