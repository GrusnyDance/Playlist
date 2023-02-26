package crud

import (
	pb "playlist/proto"
	"playlist/server/entity"
)

// Server implements gRPC server
type Server struct {
	pb.UnimplementedPlaylistServer
	PlayList *entity.Playlist
}

// NewServer is a constructor for Server
func NewServer() *Server {
	return &Server{
		PlayList: entity.NewPlaylist(),
	}
}
