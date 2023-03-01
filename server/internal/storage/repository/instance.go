package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Instance struct {
	Db *pgxpool.Pool
}

type MyTrack struct {
	Id       int
	Created  time.Time
	Name     string
	Duration int64
}

func (i *Instance) GetAll() ([]*MyTrack, error) {
	var tracks []*MyTrack
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	rows, err := i.Db.Query(ctx, "SELECT * FROM mytracks ORDER BY 1;")
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("no rows")
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		t := MyTrack{}
		rows.Scan(&t.Id, &t.Created, &t.Name, &t.Duration)
		tracks = append(tracks, &t)
	}
	return tracks, nil
}

func (i *Instance) GetTotalNum() (int, error) {
	var total int
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel()

	row := i.Db.QueryRow(ctx, "SELECT COUNT(*) FROM mytracks;")
	// Query возвращает структуру pgx.Row

	if err := row.Scan(&total); err == pgx.ErrNoRows {
		return 0, fmt.Errorf("no rows yet")
	}
	return total, nil
}

func (i *Instance) Insert(name string, duration int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel()

	_, err := i.Db.Exec(ctx, "INSERT INTO tracks (created_at, name, duration) VALUES ($1, $2, $3);",
		time.Now(), name, duration)
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) Delete(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel()

	_, err := i.Db.Exec(ctx, "DELETE FROM tracks WHERE name = $1;", name)
	if err != nil {
		return err
	}
	return nil
}
