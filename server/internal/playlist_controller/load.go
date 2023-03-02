package playlist_controller

import (
	"bufio"
	"os"
	"playlist/server/internal/server_crud"
	"strconv"
	"strings"
)

func Start(srv *server_crud.Server) error {
	if err := LoadFromDbToPlaylist(srv); err != nil {
		return err
	}
	if srv.PlayList.NumOfTracks == 0 {
		return nil
	} else {
		setCursor(srv)
	}
	return nil
}

func Finish(srv *server_crud.Server) {
	// Open file for writing, truncating if it already exists
	file, _ := os.OpenFile("./server/internal/playlist_controller/config.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()

	// Write some content to the file
	file.WriteString("current_cursor=" + srv.PlayList.CurrentPlay.Name + "\n")
	off := strconv.Itoa(int(srv.PlayList.CurrentPlay.CurrentOffset))
	file.WriteString("current_offset=" + off + "\n")
}

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

func setCursor(srv *server_crud.Server) {
	if srv.PlayList.NumOfTracks == 0 {
		return
	}

	file, _ := os.Open("./server/internal/playlist_controller/config.txt")
	defer file.Close()

	// Create a new scanner and set the split function
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Set current cursor
	scanner.Scan()
	str := strings.Split(scanner.Text(), "=")
	if str[1] == "" {
		return
	} else {
		for tmp := srv.PlayList.Tracks; tmp != nil; tmp = tmp.Next {
			if tmp.Name == str[1] {
				srv.PlayList.CurrentCursor = tmp
			}
		}
	}

	// Set current offset
	scanner.Scan()
	str = strings.Split(scanner.Text(), "=")
	if off, _ := strconv.Atoi(str[1]); off == 0 {
		return
	} else {
		srv.PlayList.CurrentCursor.CurrentOffset = int64(off)
	}
}
