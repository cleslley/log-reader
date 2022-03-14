package repository

import (
	"bufio"
	"encoding/json"
	"log"
	"log-reader/src/db"
	"log-reader/src/model"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestLogInsert(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	file, err := os.Open("log_test.txt")
	assert.NoError(t, err)
	defer file.Close()

	var log model.Log
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err := json.Unmarshal([]byte(scanner.Text()), &log)
		assert.NoError(t, err)
	}

	l := NewLogRepository(conn)
	err = l.Insert(&log)
	assert.NoError(t, err)
}
