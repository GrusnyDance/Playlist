package server_crud

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	pb "playlist/proto"
	"playlist/server/internal/playlist"
	"playlist/server/internal/storage/repository"
)

// Server implements gRPC server
type Server struct {
	pb.UnimplementedPlaylistServer
	PlayList       *playlist.Playlist
	YoutubeService *youtube.Service
	DbInstance     *repository.Instance
}

// NewServer is a constructor for Server
func NewServer() (*Server, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		return nil, err
	}

	// пул подключений к базе
	pool, err := repository.InitRep()
	if err != nil {
		log.Println(err)
	}

	instance := repository.Instance{Db: pool}

	return &Server{
		PlayList:       playlist.NewPlaylist(),
		YoutubeService: youtubeService,
		DbInstance:     &instance,
	}, nil
}
