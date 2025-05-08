// The player package implenents the player struct
// and its associated functions.
// author: Nicholas Sutton
package player

import (
	"github.com/google/uuid"
	"github.com/nick-Sutton/Gaggle/backend/internal/time"
)

// A Player stores the players first name, last name,
// email, a unique Id, and a list of the players
// available time slots.
type Player struct {
	FirstName          string
	LastName           string
	Email              string
	Id                 string
	AvailableTimeSlots []time.TimeSlot
}

// The NewPlayer function creates a player struct and returns
// a pointer to the newely created player.
// param: the players first name
// param: the players last name
// param: the players email
// return: a pointer to the new player
func NewPlayer(firstName string, lastName string, email string) *Player {
	p := &Player{}
	p.FirstName = firstName
	p.LastName = lastName
	p.Email = email
	p.Id = uuid.NewString()
	p.AvailableTimeSlots = make([]time.TimeSlot, 0, 20)

	return p
}
