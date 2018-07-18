package i2paddresshelper

import (
	"log"
	"testing"
)

func TestAddressHelper(t *testing.T) {
	addresshelper, err := NewI2pAddressHelper("http://joajgazyztfssty4w2on5oaqksz6tqoxbduy553y34mf4byv6gpq.b32.i2p", "127.0.0.1", "7656", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addresshelper.QueryHelper("i2pforum.i2p"))
}
