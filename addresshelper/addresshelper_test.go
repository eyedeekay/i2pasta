package i2paddresshelper

import (
	"log"
	"testing"
)

func TestAddressHelper(t *testing.T) {
	addresshelper, err := NewI2pAddressHelper("http://inr.i2p", "127.0.0.1", "7656")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addresshelper.QueryHelper("i2p-projekt.i2p"))
	log.Println(addresshelper.QueryHelper("i2pforum.i2p"))
}
