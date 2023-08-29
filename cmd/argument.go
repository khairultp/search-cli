package cmd

import (
	"errors"
	"flag"
	"fmt"
)

var ErrRequiredTags = errors.New("tags is required")

type Argument struct {
	tags  string
	fetch bool
}

func (a *Argument) String() string {
	return fmt.Sprintf("tags=%s", a.tags)
}

func parseArgument() (*Argument, error) {
	var argument Argument
	flag.StringVar(&argument.tags, "tags", "", "tags for specific users. separate by comma ex. -tag=sometag,anothertag")
	flag.BoolVar(&argument.fetch, "fetch", true, "fetch users data ex. -fetch=false")

	flag.Parse()

	if argument.tags == "" {
		return nil, ErrRequiredTags
	}

	return &argument, nil
}
