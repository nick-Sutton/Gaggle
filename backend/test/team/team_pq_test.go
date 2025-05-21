package team_test

import (
	"testing"

	"github.com/nick-Sutton/Gaggle/backend/internal/team"
)

func TestNewTeapPQ(t *testing.T) {
	tpq := team.NewTeamPQ()

	tpq = team.RebuildPQ(tpq)
}
