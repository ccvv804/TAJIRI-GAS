package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/rasky/go-lzo"
)

func tgas(indata []byte) (outdata []byte) {
	dat, err := lzo.Decompress1X(bytes.NewReader(indata), len(indata), 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	midisize := binary.LittleEndian.Uint32(dat[32:36])
	outdata = dat[52 : 52+midisize]
	return
}

func main() {
	file := flag.String("file", "00000.tjr", "Input file")
	flag.Parse()
	dat, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("TAJIRJ GAS V1 (public)")
	if dat[0] == byte(69) && dat[1] == byte(84) && dat[2] == byte(74) && dat[3] == byte(78) {
		fmt.Println("ETJN!")
	} else {
		if dat != nil {
			err := ioutil.WriteFile(*file+".mid", tgas(dat), 0755)
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Println("OK")
			}
		}
	}
}
