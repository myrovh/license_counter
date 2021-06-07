package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"encoding/csv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type computer struct {
	computerId   string
	computerType ComputerType
}

//go:generate stringer -type=ComputerType

type ComputerType int

const (
	LAPTOP ComputerType = iota
	DESKTOP
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	fileName := flag.String("file", "", "filename of the csv file to parse")
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal().Err(err).Msg("failed while calculating licenses")
	}

	total, err := CalculateTotalLicenses(bufio.NewReader(file), "374")
	if err != nil {
		log.Fatal().Err(err).Msg("failed while calculating licenses")
	}

	fmt.Println(total)
}

func CalculateTotalLicenses(reader io.Reader, application string) (totalLicenses int, err error) {
	users, err := parseCsv(reader, application)

	for _, user := range users {
		var countDesktop int
		var countLaptop int
		var userLicenses int

		for _, computer := range user {
			switch computer.computerType {
			case DESKTOP:
				countDesktop++
			case LAPTOP:
				countLaptop++
			}
		}

		if countDesktop < countLaptop {
			userLicenses = countLaptop
		} else {
			userLicenses = countDesktop
		}

		log.Debug().
			Int("desktop_count", countDesktop).
			Int("laptop_count", countLaptop).
			Msg("debug out")

		totalLicenses = totalLicenses + userLicenses
	}

	return
}

func parseCsv(reader io.Reader, applicationID string) (users map[string][]computer, err error) {
	csvReader := csv.NewReader(reader)

	users = make(map[string][]computer)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return users, err
		}

		if row[2] == applicationID {
			var newComputer computer
			var duplicateFound bool

			for _, computer := range users[row[1]] {
				if computer.computerId == row[0] {
					log.Debug().
						Str("existing_id", computer.computerId).
						Str("existing_type", computer.computerType.String()).
						Strs("row_contents", row).
						Msg("duplicate machine id found: skipping")
					duplicateFound = true
				}
			}

			if duplicateFound {
				continue
			}

			inputComputerType := strings.ToUpper(row[3])
			newComputer.computerId = row[0]

			if inputComputerType == DESKTOP.String() {
				newComputer.computerType = DESKTOP
			} else if inputComputerType == LAPTOP.String() {
				newComputer.computerType = LAPTOP
			} else {
				log.Warn().Strs("row_contents", row).Msg("row isn't  or LAPTOP: skipping")
				continue
			}

			users[row[1]] = append(users[row[1]], newComputer)
		}

	}

	return
}
