package main

import (
	"fmt"
	"log"
	"nba-simulation/pkg/config"
)

func main() {
	appConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}

	fmt.Println("Total Game Duration:", appConfig.TotalGameDuration)
	fmt.Println("Game Duration Per Min", appConfig.GameDurationPerMin)
	fmt.Println("Max Attack Duration:", appConfig.MaxAttackDuration)
}
