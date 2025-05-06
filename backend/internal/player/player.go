package player

import "github.com/nick-Sutton/Gaggle/backend/internal/time"

type Player struct {
	FirstName          string
	LastName           string
	Email              string
	AvailableTimeSlots map[int]time.DailyTimeSlot // day -> availability
	PreferedTimeSlots  []time.DailyTimeSlot
}

func NewPlayer(name string) *Player {
	return &Player{FirstName: name,
		AvailableTimeSlots: make(map[int]time.DailyTimeSlot)}
}
