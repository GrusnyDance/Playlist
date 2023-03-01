package storage_playlist_controller

import (
	"fmt"
	"playlist/server/internal/server_crud"
)

func LoadFromDbToPlaylist(srv *server_crud.Server) error {
	tracks, err := srv.DbInstance.GetAll()
	if err == fmt.Errorf("no rows") {
		return nil
	} else if err != nil {
		return err
	}

	for _, val := range tracks {
		srv.PlayList.Add(val.Name, int(val.Duration))
	}
	total, err := srv.DbInstance.GetTotalNum()
	if err != nil {
		return err
	}
	srv.PlayList.NumOfTracks = uint(total)
	return nil
}
