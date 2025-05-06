package player

import (
	"github.com/google/uuid"
	"github.com/nick-Sutton/Gaggle/backend/internal/time"
)

type Player struct {
	FirstName          string
	LastName           string
	Email              string
	Id                 string
	AvailableTimeSlots map[int]time.DailyTimeSlot // day -> availability
	PreferedTimeSlots  []time.DailyTimeSlot
}

func NewPlayer(firstName string, lastName string, email string) *Player {
	p := &Player{}
	p.FirstName = firstName
	p.LastName = lastName
	p.Email = email
	p.Id = uuid.NewString()
	p.AvailableTimeSlots = make(map[int]time.DailyTimeSlot)

	return p
}
