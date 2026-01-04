package io

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/nick-Sutton/Gaggle/backend/internal/player"
)

func parsePlayerData(inputFile *os.File,
	lowSkillBucket *player.PlayerPQ, medSkillBucket player.PlayerPQ,
	highSkillBucket player.PlayerPQ, weekdayCount *[7]int) {

	r := csv.NewReader(inputFile)

	for {
		// Get Line and err value
		line, err := r.Read()
		if err == io.EOF {
			break
		}

		// Check Error Value
		if err != nil {
			log.Fatal(err)
		}

		// Get the current players information
		firstName := line[1]
		lastName := line[2]
		email := line[3]

		// Add logic for getting days here
		

		// Create a player struct for the the current player
		p := player.NewPlayer(firstName, lastName, email)
		p.AvailableTimeSlots

		// Set Player skill rating and add them to the corresponding PQ
		switch strings.ToLower(line[4]) {
		case "novice":
			p.Skill = player.Novice
			lowSkillBucket.Push(p)

		case "intermediate":
			p.Skill = player.Intermediate
			medSkillBucket.Push(p)

		case "experienced":
			p.Skill = player.Experienced
			highSkillBucket.Push(p)
		default:
			log.Fatal("Skill Rating Does Not Exist")
		}

	}
}

func WriteCSV() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	file, _ := os.Create("Teams.csv")
	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing to file:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
