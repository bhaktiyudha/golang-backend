package model

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Logger struct {
	ID         string
	Method     string
	ClientIP   string
	Latency    float64
	Path       string
	StatusCode int
	TimeStamp  string
}

type LogModel struct {
	DB *sql.DB
}

func (lm LogModel) InsertLog(log *Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := "INSERT INTO users (id, clientip, latency, method, path, statuscode, timestamp) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := lm.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, log.ID, log.ClientIP, log.Latency, log.Method, log.Path, log.StatusCode, log.TimeStamp)
	return err
}
