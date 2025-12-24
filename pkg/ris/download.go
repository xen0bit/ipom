package ris

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
)

var (
	RISWhoisv4Url = "https://www.ris.ripe.net/dumps/riswhoisdump.IPv4.gz"
	RISWhoisv6Url = "https://www.ris.ripe.net/dumps/riswhoisdump.IPv6.gz"
)

func download(url string) (fileBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return fileBytes, err
	}
	defer resp.Body.Close()
	fileBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return fileBytes, err
	} else {
		return fileBytes, nil
	}
}

func RISWhoisV4() (rw string, err error) {
	fb, err := download(RISWhoisv4Url)
	if err != nil {
		return "", err
	} else {
		reader := bytes.NewReader(fb)
		gzreader, err := gzip.NewReader(reader)
		if err != nil {
			return "", err
		} else {
			fb, err = io.ReadAll(gzreader)
			if err != nil {
				return "", err
			} else {
				return string(fb), nil
			}
		}
	}
}

func RISWhoisV6() (rw string, err error) {
	fb, err := download(RISWhoisv6Url)
	if err != nil {
		return "", err
	} else {
		reader := bytes.NewReader(fb)
		gzreader, err := gzip.NewReader(reader)
		if err != nil {
			return "", err
		} else {
			fb, err = io.ReadAll(gzreader)
			if err != nil {
				return "", err
			} else {
				return string(fb), nil
			}
		}
	}
}
