package playlist

func (p *Playlist) Add(name string, duration int) {
	p.Lock()
	defer p.Unlock()

	track := &Track{
		Next: nil,
		Prev: nil,
	}

	track.Duration = duration
	track.Name = name

	p.NumOfTracks++
	if p.NumOfTracks > 1 {
		track.Prev = p.LastTrack
		p.LastTrack.Next = track
	} else {
		p.Tracks = track
		p.CurrentCursor = track
	}
	p.LastTrack = track
}
