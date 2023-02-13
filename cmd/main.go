package main

import (
	"fmt"

	"github.com/pczyk/chessapi/chessapi"
)

func main() {
	api := chessapi.NewChessAPI()
	hikaru, _ := api.FindPlayerStatsByUsername("hikaru")
	magnus, _ := api.FindPlayerStatsByUsername("MagnusCarlsen")

	if hikaru.BlitzStats.Last.Rating > magnus.BlitzStats.Last.Rating {
		fmt.Println("Hikaru")
	} else {
		fmt.Println("Magnus")
	}

	// player, _ := api.FindPlayerByUsername("hikaru")
	// fmt.Printf("%#v\n", player)
	// playerStats, _ := api.FindPlayerStatsByUsername("hikaru")
	// fmt.Printf("%#v\n", playerStats)
}
