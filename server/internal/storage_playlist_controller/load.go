package storage_playlist_controller

import (
	"playlist/server/internal/server_crud"
)

func LoadFromDbToPlaylist(srv *server_crud.Server) error {
	tracks, isNoRows, err := srv.DbInstance.GetAll()
	if isNoRows {
		return nil
	} else if err != nil {
		return err
	}

	for _, val := range tracks {
		srv.PlayList.Add(val.Name, int(val.Duration))
	}
	total := srv.DbInstance.GetTotalNum()
	srv.PlayList.NumOfTracks = total
	return nil
}
