package app

import (
	"fmt"
	"github.com/buts00/Graph/internal/database"
)

func PrintNodes(db *database.PostgresDB) error {
	nodes, err := database.Nodes(db)
	if err != nil {
		return err
	}
	fmt.Println("------Nodes------")
	for _, node := range *nodes {
		fmt.Printf("Node value %d \n", node.Value)
	}

	return nil
}

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
