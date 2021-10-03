package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pedroalbanese/gost-shred"
)

var (
	iter   = flag.Int("i", 25, "Iterations.")
	remove = flag.Bool("r", false, "Remove file(s) afterwards.")
	target = flag.String("t", "", "Target file, directory or wildcard.")
)

func main() {
	flag.Parse()

	if target == "" {
		fmt.Printf("GOST R 50739-95 Data Sanitization Method - ALBANESE Lab (c) 2020-2021\n\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("%s [-r] [-i N] -t <file.ext>\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	shredder := shred.Shredder{}
	shredconf := shred.NewShredderConf(&shredder, shred.WriteZeros|shred.WriteRand, Iter, Remove)
	matches, err := filepath.Glob(target)
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		err := shredconf.ShredDir(match)
		if err != nil {
			log.Fatal(err)
		}
	}
}
