package main

import (
	"github.com/rs/zerolog"
	"io"
	"os"
	"strings"
	"testing"
)

const test1 = `
ComputerID,UserID,ApplicationID,ComputerType,Comment
1,1,374,LAPTOP,Exported from System A
2,1,374,DESKTOP,Exported from System A`

const test2 = `
ComputerID,UserID,ApplicationID,ComputerType,Comment
1,1,374,LAPTOP,Exported from System A
2,1,374,DESKTOP,Exported from System A
3,2,374,DESKTOP,Exported from System A
4,2,374,DESKTOP,Exported from System A`

const test3 = `
ComputerID,UserID,ApplicationID,ComputerType,Comment
1,1,374,LAPTOP,Exported from System A
2,2,374,DESKTOP,Exported from System A
2,2,374,desktop,Exported from System B`

func TestCsvInputs(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	file, err := os.Open("./example_files/sample-small.csv")
	if err != nil {
		panic(err)
	}

	file2, err := os.Open("./example_files/sample-large.csv")
	if err != nil {
		panic(err)
	}

	var tests = []struct {
		name          string
		input         io.Reader
		applicationID string
		want          int
	}{
		{"readme test 1", strings.NewReader(test1), "374", 1},
		{"readme test 2", strings.NewReader(test2), "374", 3},
		{"readme test 3", strings.NewReader(test3), "374", 2},
		{"test small sample file", file, "374", 190},
		{"test large sample file", file2, "374", 15336},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans, err := CalculateTotalLicenses(test.input, test.applicationID)
			if err != nil {
				t.Errorf("got error while calculating: %s", err)
			}
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
