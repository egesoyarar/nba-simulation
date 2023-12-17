package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"nba-simulation/pkg/config"
	"nba-simulation/pkg/entities"
	"net/http"
	"sync"
	"time"
)

var playerChoices = []int{0, 1, 2, 3, 4}

var (
	homeScore                                                                                 int
	awayScore                                                                                 int
	player1, player2, player3, player4, player5, player6, player7, player8, player9, player10 int
	mu                                                                                        sync.Mutex
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
	scoresCopy := [12]int{homeScore, awayScore, player1, player2, player3, player4, player5, player6, player7, player8, player9, player10}
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

	if points != 0 {
		scorerOrder, assistOrder := randomPickScorerAssist()
		team.Points += points
		team.Players[scorerOrder].Points += points
		team.Players[assistOrder].Assists += 1
		fmt.Printf("%v scores %v points and assisted by %v\n", team.Players[scorerOrder].Name, points, team.Players[assistOrder].Name)
	}

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
		player1 = homeTeam.Players[0].Points
		player2 = homeTeam.Players[1].Points
		player3 = homeTeam.Players[2].Points
		player4 = homeTeam.Players[3].Points
		player5 = homeTeam.Players[4].Points
		player6 = awayTeam.Players[0].Points
		player7 = awayTeam.Players[1].Points
		player8 = awayTeam.Players[2].Points
		player9 = awayTeam.Players[3].Points
		player10 = awayTeam.Players[4].Points
		mu.Unlock()
	}
}

func randomPickScorerAssist() (int, int) {
	scorer := rand.Intn(5)
	restPlayerChoices := append(playerChoices[:scorer], playerChoices[scorer+1:]...)
	assist := rand.Intn(len(restPlayerChoices))
	return scorer, assist
}
