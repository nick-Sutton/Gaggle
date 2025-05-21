// Provides the implementation for a priority queue of teams
// author: Nicholas Sutton
package team

import (
	"github.com/nick-Sutton/Gaggle/backend/internal/pq"
)

// Dynamic Heap for integers
type TeamPQ struct {
	*pq.PriorityQueue[Team]
}

// Create a PQ containing teams and defines
// the PQ's comparison function.
// When a TeamPQ is is a max heap.
func NewTeamPQ() *TeamPQ {
	maxCmp := func(a, b Team) bool {
		aLen := len(a.TeamMembers)
		bLen := len(b.TeamMembers)
		if aLen != bLen {
			return aLen > bLen
		}
		return true
	}

	return &TeamPQ{pq.NewPriorityQueue(maxCmp)}
}

// Rebuilds the a maxheap TeamPQ into a minHeap
func RebuildPQ(maxTeamPQ *TeamPQ) *TeamPQ {
	minCmp := func(a, b Team) bool {
		aLen := len(a.TeamMembers)
		bLen := len(b.TeamMembers)
		if aLen != bLen {
			return aLen < bLen
		}
		return true
	}

	teamSlice := maxTeamPQ.GetItems()
	minTeamPQ := TeamPQ{pq.NewPriorityQueueFromSlice(teamSlice, minCmp)}

	return &minTeamPQ

}
