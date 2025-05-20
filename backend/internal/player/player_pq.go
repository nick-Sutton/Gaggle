// defines a priority queue consisting of Players
// author: Nora Cam
package player

import (
	"github.com/nick-Sutton/Gaggle/backend/internal/pq"
)

// Min-heap for integers
type PlayerPQ struct {
	*pq.PriorityQueue[Player]
}

func NewPlayerPQ() *PlayerPQ {
	cmp := func(a, b Player) bool {
		aLen := len(a.AvailableTimeSlots)
		bLen := len(b.AvailableTimeSlots)
		if aLen != bLen {
			return aLen < bLen
		}
		return true
	}
	return &PlayerPQ{pq.NewPriorityQueue(cmp)}
}
