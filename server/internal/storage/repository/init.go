package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	_ "playlist/server/internal/storage/migrations"
)

func InitRep() (*pgxpool.Pool, error) {
	// создаем конфиг
	poolConfig, err := NewPoolConfig()
	if err != nil {
		return nil, fmt.Errorf("Pool config error: %v\n", err)
	}
	// Макс количество соединений, которые могут находиться в ожидании
	poolConfig.MaxConns = 10

	// Создаем пул подключений
	pool, err := NewConnection(poolConfig)
	if err != nil {
		return nil, fmt.Errorf("Connect to database failed: %v\n", err)
	}

	// Проверяем подключение
	_, err = pool.Exec(context.Background(), ";")
	if err != nil {
		return nil, fmt.Errorf("Ping failed: %v\n", err)
	}

	//применить миграции через стандартный драйвер database/sql
	mdb, err := sql.Open("postgres", poolConfig.ConnString())
	err = mdb.Ping()
	err = goose.Up(mdb, "./server/internal/storage/migrations")
	if err != nil {
		panic(err)
	}
	mdb.Close()

	return pool, nil
}

// обертка для создания подключения с помощью пула
func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
