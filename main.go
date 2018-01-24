package main

import (
	"github.com/hackathon/journeys/cmd"

	// register protocols and resources
	_ "github.com/hackathon/journeys/protocols"
	_ "github.com/hackathon/journeys/resources"
)

func main() {
	cmd.RootCmd.Execute()
}
