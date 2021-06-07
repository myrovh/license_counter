package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/rs/zerolog"
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

const test4 = `
ComputerID,UserID,ApplicationID,ComputerType,Comment,extra
1,1,374,LAPTOP,Exported from System A,te
2,2,374,DESKTOP,Exported from System A,te
2,2,374,desktop,Exported from System B,te`

const test5 = `not csv`

type tests struct {
	name          string
	input         io.Reader
	applicationID string
	want          int
	pass          bool
}

func TestMain(m *testing.M) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	os.Exit(m.Run())
}

func BenchmarkCsv(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	file2, err := os.Open("./example_files/sample-large.csv")
	if err != nil {
		b.Skip("large file not found skipping")
	}

	var test = tests{"large sample file", file2, "374", 15336, true}


	_, err = CalculateTotalLicenses(test.input, test.applicationID)
	if err != err {
		b.Errorf("failed while calculating: %s", err)
	}
}

func TestCsvInputs(t *testing.T) {
	file, err := os.Open("./example_files/sample-small.csv")
	if err != nil {
		panic(err)
	}

	var tests = []tests{
		{"readme test 1", strings.NewReader(test1), "374", 1, true},
		{"readme test 2", strings.NewReader(test2), "374", 3, true},
		{"readme test 3", strings.NewReader(test3), "374", 2, true},
		{"small sample file", file, "374", 190, true},
		{"row count", strings.NewReader(test4), "374", 0, false},
		{"not csv", strings.NewReader(test5), "374", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans, err := CalculateTotalLicenses(test.input, test.applicationID)
			if err != nil {
				if test.pass {
					t.Errorf("got error while calculating: %s", err)
				}
			}
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
