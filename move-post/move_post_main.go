package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h bool   // need help
	s string // source direction
	t string // target direction
)

func init() {
	flag.StringVar(&s, "s", "", "source direction")
	flag.StringVar(&t, "t", "", "target direction")
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `Options:
	-h help
	-s source direction
	-t target direction`)
}

func main() {
	flag.Parse()
	if h {
		usage()
	}

	if len(s) > 0 && len(t) > 0 {
		err := move(s, t)
		checkErr(err)
	} else {
		println("ERROR: -s & -t can not be null")
	}
}
