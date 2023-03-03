package playlist_controller

import (
	"bufio"
	"fmt"
	"os"
	"playlist/server/internal/server_crud"
	"strconv"
	"strings"
)

func Start(srv *server_crud.Server) error {
	srv.PlayList.Add("Валерий Меладзе - Иностранец", 248)
	srv.DbInstance.Insert("Валерий Меладзе - Иностранец", 248)

	if err := LoadFromDbToPlaylist(srv); err != nil {
		fmt.Println(err)
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
	fmt.Println("I finish")
	f, _ := os.Create("./server/internal/playlist_controller/config.txt")
	defer f.Close()

	// Write some content to the file
	if srv.PlayList.CurrentPlay == nil {
		return
	}
	f.WriteString("current_cursor=" + srv.PlayList.CurrentPlay.Name + "\n")
	off := strconv.Itoa(int(srv.PlayList.CurrentPlay.CurrentOffset))
	f.WriteString("current_offset=" + off + "\n")
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

	filename := "./server/internal/playlist_controller/config.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
		return
	}

	file, _ := os.Open(filename)
	defer file.Close()
	fi, _ := file.Stat()
	if fi.Size() == 0 {
		return
	}

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
