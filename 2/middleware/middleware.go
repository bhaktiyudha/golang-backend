package middleware

import (
	"database/sql"
	"log"
	"mission-2/config"
	"mission-2/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CORS Middleware
func CORS(c *gin.Context) {
	// Headers Set
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	//Handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func DBConn(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("sql database", db)
		c.Next()
	}
}

func LoggerMiddleware(params gin.LogFormatterParams) string {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error run API : %s", err)
	}
	model.LogModel{
		DB: db,
	}.InsertLog(&model.Logger{
		ID:         uuid.New().String(),
		Method:     params.Method,
		ClientIP:   string(params.ClientIP),
		Latency:    float64(params.Latency.Seconds()),
		Path:       params.Path,
		StatusCode: int(params.StatusCode),
		TimeStamp:  params.TimeStamp.Format("2006-01-02 15:04:05"),
	})
	return ""
}
