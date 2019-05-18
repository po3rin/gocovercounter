package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/po3rin/gocovercounter"
)

var (
	output = flag.String("o", "", "file for output; default: stdout")
)

func main() {
	var w io.Writer
	flag.Parse()

	if *output != "" {
		var err error
		w, err = os.Create(*output)
		if err != nil {
			log.Fatalf("coco: %s", err)
		}
	} else {
		w = os.Stdout
	}

	if flag.Arg(0) == "" {
		log.Fatalf("coco: require target file path")
	}

	err := gocovercounter.Annotate(w, flag.Arg(0))
	if err != nil {
		log.Fatalf("coco: %s", err)
	}
}
