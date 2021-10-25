// GOST-Shred -- GOST R 50739-95 Data Sanitization Method
// Copyright (C) 2020-2021 Pedro Albanese <pedroalbanese@hotmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the ISC Public License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// Command-line GOST R 50739-95 shred function.
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
	iter   = flag.Int("i", 3, "Iterations.")
	remove = flag.Bool("r", false, "Remove file(s) afterwards.")
	target = flag.String("t", "", "Target file, directory or wildcard.")
)

func main() {
	flag.Parse()

	if *target == "" {
		fmt.Printf("GOST R 50739-95 Data Sanitization Method - ALBANESE Lab (c) 2020-2021\n\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("%s [-r] [-i N] -t <file.ext>\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	shredder := shred.Shredder{}
	shredconf := shred.NewShredderConf(&shredder, shred.WriteZeros|shred.WriteRand, *iter, *remove)
	matches, err := filepath.Glob(*target)
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
