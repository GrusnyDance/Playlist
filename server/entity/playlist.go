package entity

import "time"

type Track struct {
	Duration      time.Duration
	Name          string
	CurrentOffset uint64
	Next          *Track
	Prev          *Track
}

type Playlist struct {
	Tracks      []Track
	LastTrack   *Track
	NumOfTracks uint
}

func NewPlaylist() *Playlist {
	return &Playlist{
		NumOfTracks: 0,
	}
}
