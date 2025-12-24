package ris

import (
	"net/netip"
)

func (rwrs *RISWhoisRecords) GetRecords(origin int) []netip.Prefix {
	var cidrs []netip.Prefix
	for _, r := range *rwrs {
		if r.Origin == origin {
			cidrs = append(cidrs, r.Prefix)
		}
	}
	return cidrs
}
