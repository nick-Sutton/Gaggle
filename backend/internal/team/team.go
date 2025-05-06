package team

import (
	"github.com/nick-Sutton/Gaggle/backend/internal/player"
	"github.com/nick-Sutton/Gaggle/backend/internal/time"
)

type Team struct {
	Name        string
	TeamMembers map[string]player.Player
	TimeSlot    time.DailyTimeSlot
}
