package main

import (
	"github.com/dbalan/cryptopals/set2"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./rc.ppm")
	if err != nil {
		panic(err)
	}
	key := []byte("YELLOW SUBMARINE")
	enc, err := set2.EncAES128ECB(data, key)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./enc.ppm", enc, 0755)
	if err != nil {
		panic(err)
	}
}
