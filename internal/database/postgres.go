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

func Edges(db *PostgresDB) (graph.Graph, error) {
	row, err := db.DB.Query("SELECT * FROM edges")
	if err != nil {
		return graph.Graph{}, err
	}

	var cur graph.Graph
	for row.Next() {
		var id, source, destination, weight int
		if err := row.Scan(&id, &source, &destination, &weight); err != nil {
			return graph.Graph{}, err
		}
		cur.Edges = append(cur.Edges, graph.Edge{Id: id, Source: source, Destination: destination, Weight: weight})
	}

	return cur, nil
}

func AddEdge(db *PostgresDB, edge graph.Edge) error {
	if _, err := db.DB.Exec("INSERT INTO edges (source_node_id, destination_node_id, weight) VALUES ($1, $2, $3)",
		edge.Source, edge.Destination, edge.Weight); err != nil {
		return err
	}
	return nil
}

func DeleteEdge(db *PostgresDB, edge graph.Edge) error {
	if _, err := db.DB.Exec("DELETE FROM edges WHERE source_node_id = $1 AND destination_node_id = $2  AND weight = $3;",
		edge.Source, edge.Destination, edge.Weight); err != nil {
		return err
	}
	return nil
}

func IsEdgeExist(db *PostgresDB, edge graph.Edge) (bool, error) {
	var count int
	if err := db.DB.QueryRow("SELECT COUNT(*) FROM edges WHERE source_node_id = $1 AND destination_node_id = $2 AND weight = $3;",
		edge.Source, edge.Destination, edge.Weight).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
