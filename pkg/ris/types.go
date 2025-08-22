package ris

import "net/netip"

type RISWhoisRecord struct {
	Origin int
	Prefix netip.Prefix
}

type RISWhoisRecords []RISWhoisRecord
