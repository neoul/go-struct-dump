package main

import (
	"fmt"

	"github.com/neoul/gostruct-dump/dump"
)

type inside struct {
	Integerval int
}

type gostruct struct {
	Integerval     int
	Integerptr     *int
	Stringval      string
	Stringptr      *string
	boolNotPresent bool
	BoolPresent    bool
	Inside         inside
}

func main() {
	i := 100
	s := "stringValue"
	gs := gostruct{
		Integerval:     10,
		Integerptr:     &i,
		Stringval:      "gostruct string",
		Stringptr:      &s,
		BoolPresent:    true,
		boolNotPresent: true,
		Inside:         inside{Integerval: 20},
	}
	// print gs to Stdout
	dump.Print(gs)

	// = dump.Print(gs)
	dump.ValueDump(gs, 3, func(x ...interface{}) { fmt.Print(x...) })
}
