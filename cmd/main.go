package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"nba-simulation/pkg/config"
	"nba-simulation/pkg/entities"
	"net/http"
	"sync"
	"time"
)

var (
	homeScore int
	awayScore int
	mu        sync.Mutex
)

func main() {
	appConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}

	totalGameDuration := appConfig.TotalGameDuration

	homeTeam := &entities.Team{
		Name: "Home Team",
		Players: []entities.Player{
			{Name: "Player 1", Points: 0, Assists: 0},
			{Name: "Player 2", Points: 0, Assists: 0},
			{Name: "Player 3", Points: 0, Assists: 0},
			{Name: "Player 4", Points: 0, Assists: 0},
			{Name: "Player 5", Points: 0, Assists: 0},
		},
		Points: 0,
	}

	awayTeam := &entities.Team{
		Name: "Away Team",
		Players: []entities.Player{
			{Name: "Player 6", Points: 0, Assists: 0},
			{Name: "Player 7", Points: 0, Assists: 0},
			{Name: "Player 8", Points: 0, Assists: 0},
			{Name: "Player 9", Points: 0, Assists: 0},
			{Name: "Player 10", Points: 0, Assists: 0},
		},
		Points: 0,
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/getScores", getScoresHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	go updateScores(homeTeam, awayTeam, totalGameDuration)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/index.html")
}

func getScoresHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Return the current scores as JSON
	scoresCopy := [2]int{homeScore, awayScore}
	scoresJSON, err := json.Marshal(scoresCopy)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(scoresJSON)
}

func simulateAttack(team *entities.Team) {
	points := randomScoring()

	team.Points += points
}

func randomScoring() int {
	choices := []int{0, 2, 3}
	i := rand.Intn(len(choices))
	return choices[i]
}

func updateScores(homeTeam, awayTeam *entities.Team, totalGameDuration int) {
	for i := 0; i < totalGameDuration; i++ {
		time.Sleep(2 * time.Second)

		mu.Lock()
		simulateAttack(homeTeam)
		simulateAttack(awayTeam)
		homeScore = homeTeam.Points
		awayScore = awayTeam.Points
		mu.Unlock()
	}
}
