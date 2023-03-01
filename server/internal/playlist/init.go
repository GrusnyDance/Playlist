package playlist

import "sync"

type Track struct {
	Duration      int
	Name          string
	CurrentOffset int64
	Next          *Track
	Prev          *Track
}

type Playlist struct {
	Tracks        *Track
	LastTrack     *Track
	NumOfTracks   int
	CurrentCursor *Track
	CurrentPlay   *Track
	sync.Mutex
	IsPlayed bool
}

func NewPlaylist() *Playlist {
	return &Playlist{
		NumOfTracks: 0,
	}
}
