package main

import (
	"bufio"
	"encoding/json"
	"log"
	"log-reader/src/db"
	"log-reader/src/model"
	"log-reader/src/report"
	"log-reader/src/repository"
	"math"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const (
	logFile = "logs.txt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(logFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	conn, err := startDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Banco de Dados Conectado com Sucesso.")

	err = dropCollection(conn)
	if err != nil && err.Error() != "ns not found" {
		log.Fatal(err)
	}

	log.Printf("Aguarde, Inserindo Logs no Banco de Dados...")
	err = insertLogsInDB(file, conn)
	if err != nil {
		log.Fatal(err)
	}

	logs, err := getLogsFromDB(conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Aguarde, Gerando Relatórios...")
	err = report.CreateConsumerReport(logs)
	if err != nil {
		log.Fatal(err)
	}

	err = report.CreateServiceReport(logs)
	if err != nil {
		log.Fatal(err)
	}

	err = report.CreateServiceTimeReport(logs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Relatórios Gerados com Sucesso.")
}

func startDbConnection() (db.Connection, error) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		return conn, err
	}
	return conn, nil
}

func insertLogsInDB(file *os.File, conn db.Connection) error {
	var log model.Log
	scanner := bufio.NewScanner(file)
	var logs []model.Log

	for scanner.Scan() {
		err := json.Unmarshal([]byte(scanner.Text()), &log)
		if err != nil {
			panic(err)
		}
		logs = append(logs, log)
	}

	splitLogs := splitArray(logs, len(logs)/4)
	var waitGroup sync.WaitGroup
	logsRepository := repository.NewLogRepository(conn)

	for i := range splitLogs {
		waitGroup.Add(1)
		go func(i int, waitGroup *sync.WaitGroup) {
			defer waitGroup.Done()
			for j := range splitLogs[i] {
				logsRepository.Insert(&splitLogs[i][j])
			}
		}(i, &waitGroup)
	}
	waitGroup.Wait()

	err := scanner.Err()
	if err != nil {
		return err
	}

	return nil
}

func getLogsFromDB(conn db.Connection) (logs []*model.Log, err error) {
	logsRepository := repository.NewLogRepository(conn)
	logs, err = logsRepository.GetAll()
	if err != nil {
		return logs, err
	}

	return logs, nil
}

func dropCollection(conn db.Connection) error {
	logsRepository := repository.NewLogRepository(conn)
	err := logsRepository.DeleteAll()
	if err != nil {
		return err
	}

	return nil
}

func splitArray(s []model.Log, size int) [][]model.Log {
	var n [][]model.Log
	if size < 1 {
		return n
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}
