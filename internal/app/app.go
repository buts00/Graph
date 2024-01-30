package app

import (
	"fmt"
	"github.com/buts00/Graph/internal/database"
)

func PrintEdges(db *database.PostgresDB) error {
	edges, err := database.Edges(db)
	if err != nil {
		return err
	}
	fmt.Println("------Edges------")
	for _, edge := range *edges {
		fmt.Printf("From %d to %d Weight = %d \n", edge.Source, edge.Destination, edge.Weight)
	}
	return nil
}
