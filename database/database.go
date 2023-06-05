package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

type DB struct {
	db    *sql.DB
	mutex sync.Mutex
}

var instance *DB

func GetDBInstance() *DB {
	if instance == nil {
		var connStr = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", DATABASE_USER, DATABASE_PASSWORD, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME)
		fmt.Println(connStr)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		if err = db.Ping(); err != nil {
			panic(err) // If any error occure Panic
		}
		instance = &DB{
			db:    db,
			mutex: sync.Mutex{},
		}
	}
	create_statements, err := read_and_parse_sql_file("./database/files/create.sql")
	if err == nil {
		for _, statement := range create_statements {
			_, err := instance.db.Exec(statement)
			if err != nil {
				log.Printf("Error while Creating Tables: %s", err.Error())
			}
		}
	} else {
		log.Printf("Error while Creating Tables: %s", err.Error())
	}
	return instance
}

func (db *DB) Close() error {
	instance = nil
	return db.db.Close()
}

func (db *DB) GetAllLogs() ([]*Log, error) {
	statement_string := "SELECT * FROM logs;"
	rows, err := db.db.Query(statement_string)
	if err != nil {
		rows.Close()
		return nil, err
	}

	var logs = []*Log{}
	for rows.Next() {
		log, err := make_log(rows)
		if err != nil {
			rows.Close()
			return nil, err
		}
		logs = append(logs, log)
	}
	rows.Close()

	return logs, nil
}

func (db *DB) GetLogsFromApplicant(applicant string) ([]*Log, error) {
	statement_string := "SELECT * FROM logs WHERE applicant = $1;"
	rows, err := db.db.Query(statement_string, applicant)
	if err != nil {
		rows.Close()
		return nil, err
	}

	var logs = []*Log{}
	for rows.Next() {
		log, err := make_log(rows)
		if err != nil {
			rows.Close()
			return nil, err
		}
		logs = append(logs, log)
	}
	rows.Close()

	return logs, nil
}

func (db *DB) WriteLog(log *Log) error {
	statement_string := "INSERT INTO logs(applicant, log_time, log_message) VALUES ($1, NOW(), $2);"
	_, err := db.db.Exec(statement_string, log.Applicant, log.Message)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func make_log(rows *sql.Rows) (*Log, error) {
	var id int
	var applicant string
	var timestamp string
	var message string

	err := rows.Scan(&id, &applicant, &timestamp, &message)

	if err != nil {
		return nil, err
	}

	return &Log{
		Id: id,
		Applicant: applicant,
		Timestamp: timestamp,
		Message: message,
	}, nil
}

func read_and_parse_sql_file(filepath string) ([]string, error) {
	var statements []string

	// Open File
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Scanner zum Zeilenweisen Lesen der Datei erstellen
	scanner := bufio.NewScanner(file)

	// Variable zum Zwischenspeichern von mehrzeiligen Statements
	var statementBuilder strings.Builder

	// Zeilenweise Datei lesen
	for scanner.Scan() {
		line := scanner.Text()

		// Wenn die Zeile mit einem Kommentar beginnt, überspringen
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			continue
		}

		// Wenn die Zeile ein Teil eines mehrzeiligen Statements ist,
		// an den Builder anhängen
		if strings.HasSuffix(strings.TrimSpace(line), ";") && statementBuilder.Len() > 0 {
			statementBuilder.WriteString(" ")
			statementBuilder.WriteString(strings.TrimSpace(line))
			statement := statementBuilder.String()
			statements = append(statements, statement)
			statementBuilder.Reset()
		} else {
			// Ansonsten die Zeile an den Builder anhängen
			statementBuilder.WriteString(" ")
			statementBuilder.WriteString(strings.TrimSpace(line))
		}
	}

	// Letztes Statement hinzufügen, falls mehrzeiliges Statement am Ende
	if statementBuilder.Len() > 0 {
		statement := statementBuilder.String()
		statements = append(statements, strings.TrimSpace(statement))
	}

	// Fehler beim Scanner prüfen
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return statements, nil
}
