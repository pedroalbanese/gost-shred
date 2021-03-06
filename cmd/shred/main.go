package main

import (
	"flag"
	"fmt"
	"github.com/pedroalbanese/gost-shred"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var Iter int
	flag.IntVar(&Iter, "i", 1, "Iterations.")
	var Remove bool
	flag.BoolVar(&Remove, "r", false, "Remove file(s) afterwards.")
	var target string
	flag.StringVar(&target, "t", "", "Target file, directory or wildcard.")
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
