package main

import (
	"flag"
	"log"
	"os"

	"github.com/po3rin/gocovercounter"
)

func main() {
	fd, err := os.Create("out.txt")
	if err != nil {
		log.Fatalf("cover: %s", err)
	}
	flag.Parse()

	gocovercounter.Annotate(fd, flag.Arg(0))
}
