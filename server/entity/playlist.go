package entity

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
	NumOfTracks   uint
	CurrentCursor *Track
	CurrentPlay   *Track
	IsPlayed      bool
}

func NewPlaylist() *Playlist {
	return &Playlist{
		NumOfTracks: 0,
	}
}
