package playlist

func (p *Playlist) Delete(deletedName string) {
	if p.CurrentPlay.Name == deletedName && p.IsPlayed {
		p.Lock()
		p.IsPlayed = false
		p.CurrentPlay = nil
		p.Unlock()
	}

	temp := p.Tracks
	for temp != nil {
		if temp.Name == deletedName {
			if p.CurrentCursor.Name == deletedName {
				if temp.Prev != nil {
					p.CurrentCursor = temp.Prev
				} else {
					p.CurrentCursor = temp.Next
				}
			}
			if temp.Prev == nil {
				p.Tracks = temp.Next
			} else if temp.Next == nil {
				p.LastTrack = temp.Prev
			} else {
				lastB4Delete := temp.Prev
				nextAfterDelete := temp.Next
				lastB4Delete.Next = nextAfterDelete
				nextAfterDelete.Prev = lastB4Delete
			}
			p.NumOfTracks--
			break
		} else {
			temp = temp.Next
		}
	}
}
