package io

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/nick-Sutton/Gaggle/backend/internal/disjoint"
	"github.com/nick-Sutton/Gaggle/backend/internal/player"
)

func ReadCSV(inputFile *os.File) {
	r := csv.NewReader(inputFile)
	df := disjoint.MakeDisjointForest()

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		firstName := line[1]
		lastName := line[2]
		email := line[3]

		// Add logic for getting days here

		p := player.NewPlayer(firstName, lastName, email)
		df.MakeSet(p)

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
