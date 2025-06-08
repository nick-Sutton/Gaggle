// Test cases for TeamPQ
// author: Nicholas Sutton
package team_test

import (
	"testing"

	"github.com/nick-Sutton/Gaggle/backend/internal/player"
	"github.com/nick-Sutton/Gaggle/backend/internal/team"
	"github.com/nick-Sutton/Gaggle/backend/internal/time"
)

func TestNewTeamPQ(t *testing.T) {
	team1 := team.NewTeam("team1", *time.NewTimeSlot(time.Monday, time.Slot_4_5pm), 3)
	team2 := team.NewTeam("team2", *time.NewTimeSlot(time.Tuesday, time.Slot_5_6pm), 3)
	team3 := team.NewTeam("team3", *time.NewTimeSlot(time.Wednesday, time.Slot_6_7pm), 3)
	team4 := team.NewTeam("team4", *time.NewTimeSlot(time.Thursday, time.Slot_8_9pm), 3)
	team5 := team.NewTeam("team5", *time.NewTimeSlot(time.Friday, time.Slot_4_5pm), 3)

	tpq := team.NewTeamPQ()
	tpq.AddTeam(team1)

	if tpq.Peek() != team1 {
		t.Errorf("AddTeam result is incorrect, got: %s, expected %s", tpq.Peek().Name, team1.Name)
	}

	tpq.AddTeam(team2)

	if tpq.Peek() != team2 {
		t.Errorf("AddTeam result is incorrect, got: %s, expected %s", tpq.Peek().Name, team2.Name)
	}
	tpq.AddTeam(team3)

	if tpq.Peek() != team3 {
		t.Errorf("AddTeam result is incorrect, got: %s, expected %s", tpq.Peek().Name, team3.Name)
	}

	tpq.AddTeam(team4)

	if tpq.Peek() != team4 {
		t.Errorf("AddTeam result is incorrect, got: %s, expected %s", tpq.Peek().Name, team4.Name)
	}

	tpq.AddTeam(team5)

	if tpq.Peek() != team5 {
		t.Errorf("AddTeam result is incorrect, got: %s, expected %s", tpq.Peek().Name, team5.Name)
	}

	p1 := player.NewPlayer("playerOneFirst", "playerOneLast", "one@email.com")
	p2 := player.NewPlayer("playerTwoFirst", "playerTwoLast", "two@email.com")
	p3 := player.NewPlayer("playerThreeFirst", "playerThreeLast", "three@email.com")
	p4 := player.NewPlayer("playerFourFirst", "playerFourLast", "four@email.com")
	p5 := player.NewPlayer("playerFiveFirst", "playerFiveLast", "five@email.com")
	p6 := player.NewPlayer("playerSixFirst", "playerSixLast", "six@email.com")
	p7 := player.NewPlayer("playerSevenFirst", "playerSevenLast", "seven@email.com")
	p8 := player.NewPlayer("playerEightFirst", "playerEightLast", "eight@email.com")
	p9 := player.NewPlayer("playerNineFirst", "playerNineLast", "nine@email.com")
	currTeam := tpq.AddToTeam(p1)
	if currTeam != team5 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team5.Name)
	}

	currTeam = tpq.AddToTeam(p2)

	if currTeam != team5 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team5.Name)
	}

	currTeam = tpq.AddToTeam(p3)

	if currTeam != team5 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team5.Name)
	}

	currTeam = tpq.AddToTeam(p4)

	if currTeam != team4 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team4.Name)
	}

	currTeam = tpq.AddToTeam(p5)

	if currTeam != team4 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team4.Name)
	}

	currTeam = tpq.AddToTeam(p6)

	if currTeam != team4 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team4.Name)
	}

	currTeam = tpq.AddToTeam(p7)

	if currTeam != team3 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team3.Name)
	}

	currTeam = tpq.AddToTeam(p8)

	if currTeam != team3 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team3.Name)
	}

	currTeam = tpq.AddToTeam(p9)

	if currTeam != team3 {
		t.Errorf("AddPlayer result is incorrect, got: %s, expected %s", currTeam.Name, team3.Name)
	}

}
