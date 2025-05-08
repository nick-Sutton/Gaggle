// The team package implenents the team struct.
// author: Nicholas Sutton
package team

import (
	"github.com/nick-Sutton/Gaggle/backend/internal/player"
	"github.com/nick-Sutton/Gaggle/backend/internal/time"
)

// A team stores its name, a map of its members, and the timeslot
// the team will meet at.
type Team struct {
	Name        string
	TeamMembers map[string]player.Player
	TimeSlot    time.TimeSlot
}

// The NewTeam function creates a team struct.
// param: the team name
// param: the teams timeslot
// return: a pointer to the new team
func NewTeam(name string, timeSlot time.TimeSlot) *Team {
	t := &Team{}
	t.Name = name
	t.TeamMembers = make(map[string]player.Player)
	t.TimeSlot = timeSlot

	return t
}
