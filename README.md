# shred
 Package shred is a golang library to mimic the functionality of the linux `shred` command, modified to fulfil GOST R 50739-95 Data Sanitization Method requisites. See https://github.com/pedroalbanese/gosttk.
 
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