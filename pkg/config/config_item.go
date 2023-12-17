package config

import (
	"time"
)

type Config struct {
	TotalGameDuration  int
	GameDurationPerMin time.Duration
	MaxAttackDuration  time.Duration
}
