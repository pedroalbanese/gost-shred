# gost-shred
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gost-shred/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gost-shred?status.png)](http://godoc.org/github.com/pedroalbanese/gost-shred)

 Package shred is a golang library to mimic the functionality of the linux `shred` command, modified to fulfil GOST R 50739-95 Data Sanitization Method requisites. See https://github.com/pedroalbanese/gosttk.

The algorithm carries out one overwriting cycle using pseudo-random numbers and protects the data from recovery by common tools. This algorithm corresponds to protection class 2 (out of 6), according to the Russian State Technical Commission classification.

## Usage
```golang
package main
import (
  "github.com/pedroalbanese/gost-shred"
)

func main(){
	shredder := shred.Shredder{}
	shredconf := shred.NewShredderConf(&shredder, shred.WriteZeros|shred.WriteRand, 1, false)
	shredconf.ShredFile("./10k")
	shredconf.ShredDir("./toShredDir")
}
```
[GoDoc](https://pkg.go.dev/github.com/pedroalbanese/gost-shred)

## License

This project is licensed under the ISC License.
