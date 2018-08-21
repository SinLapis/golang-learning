package main

import (
	"flag"
)

func cat(numFlag bool, fileName string) {

}
func main() {
	flag.Parse()
	numFlag := false
	if "-n" == flag.Arg(0) {
		numFlag = true
		fileName := flag.Arg(1)
	} else {
		fileName := flag.Arg(0)
	}
}
