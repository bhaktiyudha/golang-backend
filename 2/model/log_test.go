package model

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var l = &Logger{
	ID:         uuid.New().String(),
	Method:     "GET",
	StatusCode: 200,
	Latency:    12,
	ClientIP:   "192.168.00.1",
	Path:       "/movie",
	TimeStamp:  "2006-10-01 10:00:00",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestLogging(t *testing.T) {
	db, mock := NewMock()

	query := "INSERT INTO users \\(id, clientip, latency, method, path, statuscode, timestamp\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(l.ID, l.ClientIP, l.Latency, l.Method, l.Path, l.StatusCode, l.TimeStamp).WillReturnResult(sqlmock.NewResult(0, 1))

	err := LogModel{DB: db}.InsertLog(l)
	assert.NoError(t, err)
}

func TestLoggingError(t *testing.T) {
	db, mock := NewMock()

	query := "INSERT INTO users \\(id, clientip, latency, method, path, statuscode, timestamp\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(l.ID, l.ClientIP, l.Latency, l.Method, l.Path, l.StatusCode, l.TimeStamp).WillReturnResult(sqlmock.NewResult(0, 0))

	err := LogModel{DB: db}.InsertLog(l)
	assert.NoError(t, err)
}
