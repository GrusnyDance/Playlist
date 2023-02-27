package crud

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
	pb "playlist/proto"
	"playlist/server/entity"
)

// Server implements gRPC server
type Server struct {
	pb.UnimplementedPlaylistServer
	PlayList       *entity.Playlist
	YoutubeService *youtube.Service
}

// NewServer is a constructor for Server
func NewServer() (*Server, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		return nil, err
	}

	return &Server{
		PlayList:       entity.NewPlaylist(),
		YoutubeService: youtubeService,
	}, nil
}
