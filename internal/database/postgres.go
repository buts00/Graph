package database

import (
	"database/sql"
	"fmt"
	"github.com/buts00/Graph/internal/app/graph"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(host, port, user, password, dbName string) (*PostgresDB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{db}, nil
}

func Edges(db *PostgresDB) (*[]graph.Edge, error) {
	row, err := db.DB.Query("SELECT * FROM edges")
	if err != nil {
		return nil, err
	}
	var cur []graph.Edge

	for row.Next() {
		var id, source, destination, weight int
		if err := row.Scan(&id, &source, &destination, &weight); err != nil {
			return nil, err

		}
		cur = append(cur, graph.Edge{Source: source, Destination: destination, Weight: weight})
	}

	return &cur, nil
}
