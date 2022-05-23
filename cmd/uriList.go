package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
)

type UriList struct {
	entry []string
}

func NewUriList() UriList {
	u := make([]string, 0, 100)
	return UriList{entry: u}
}

func (u *UriList) UriListFromArgs(args []string) {
	u.entry = append(u.entry, args...)
}

func (u *UriList) UriListFromFile(uriFilename string) error {
	if _, err := os.Stat(uriFilename); err == nil {
		return u.parseFile(uriFilename)
	} else {
		return err
	}
}

func (u *UriList) parseFile(uriFilename string) error {
	f, err := os.Open(uriFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		// validate uri
		line := scanner.Text()
		_, err := url.ParseRequestURI(line)
		if err != nil {
			if len(line) == 0 {
				fmt.Println("Received an empty line")
			} else {
				fmt.Println("Invalid uri:", line)
			}
		} else {
			// Store the uri
			u.entry = append(u.entry, line)
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if len(u.entry) == 0 {
		errMsg := fmt.Sprintln("No valid uri in ", uriFilename)
		return errors.New(errMsg)
	}

	return nil
}

func (u *UriList) Dump() {
	fmt.Println("len", len(u.entry))
	for i := 0; i < len(u.entry); i++ {
		fmt.Println("list", i, u.entry[i])
	}
}
