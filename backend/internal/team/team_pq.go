// Provides the implementation for a priority queue of teams
// author: Nicholas Sutton
package team

import (
	"github.com/nick-Sutton/Gaggle/backend/internal/player"
	"github.com/nick-Sutton/Gaggle/backend/internal/pq"
)

// Dynamic Heap for integers
type TeamPQ struct {
	*pq.PriorityQueue[*Team]
}

// Create a PQ containing teams and defines
// the PQ's comparison function.
// When a TeamPQ is is a max heap.
func NewTeamPQ() *TeamPQ {
	maxCmp := func(a, b *Team) bool {
		aLen := a.NumMembers
		bLen := b.NumMembers

		if aLen == a.MaxMembers {
			return false
		}

		if bLen == b.MaxMembers {
			return true
		}

		if aLen != bLen {
			return aLen > bLen
		}

		return true
	}

	return &TeamPQ{pq.NewPriorityQueue(maxCmp)}
}

// Rebuilds the a maxheap TeamPQ into a minHeap
func RebuildPQ(maxTeamPQ *TeamPQ) *TeamPQ {
	minCmp := func(a, b *Team) bool {
		aLen := a.NumMembers
		bLen := b.NumMembers
		if aLen != bLen {
			return aLen < bLen
		}
		return true
	}

	// Consider adding items to new PQ after construction
	teamSlice := maxTeamPQ.GetItems()
	minTeamPQ := TeamPQ{pq.NewPriorityQueueFromSlice(teamSlice, minCmp)}

	return &minTeamPQ

}

// Adds a team to the PQ
func (tpq *TeamPQ) AddTeam(team *Team) {
	tpq.PriorityQueue.PushItem(team)
}

// Adds a player to the team at the front of the pq
func (tpq *TeamPQ) AddToTeam(player *player.Player) *Team {
	currTeam := tpq.PopItem()
	currTeam.AddPlayer(player)
	tpq.PushItem(currTeam)
	return currTeam
}
