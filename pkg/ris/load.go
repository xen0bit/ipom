package ris

import (
	"bufio"
	"log"
	"net/netip"
	"os"
	"strconv"
	"strings"
)

func LoadV4() (RISWhoisRecords, error) {
	var rwrs RISWhoisRecords
	file, err := os.Open("riswhoisv4.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "%") && len(line) != 0 {
			cols := strings.Split(line, "\t")
			//sometimes 123
			//sometimes {123}
			//sometimes {12302,39737,42599,48955}
			//who the hell made this trash
			mo := strings.ReplaceAll(cols[0], "{", "")
			mo = strings.ReplaceAll(mo, "}", "")
			mol := strings.Split(mo, ",")
			for _, o := range mol {
				origin, err := strconv.Atoi(o)
				if err != nil {
					log.Fatal(err)
				}
				prefix, err := netip.ParsePrefix(cols[1])
				if err != nil {
					log.Fatal(err)
				}
				rwrs = append(rwrs, RISWhoisRecord{
					Origin: origin,
					Prefix: prefix,
				})
			}
		}
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return rwrs, nil
}
