package ris

import (
	"log"
	"net/netip"
	"strconv"
)

func (rwrs *RISWhoisRecords) GetRecords(origin string) []netip.Prefix {
	var cidrs []netip.Prefix
	oi, err := strconv.Atoi(origin)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range *rwrs {
		if r.Origin == oi {
			cidrs = append(cidrs, r.Prefix)
		}
	}
	return cidrs
}
