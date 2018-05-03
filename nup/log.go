package i2pasta ///convert

import (
	"log"
	"os"
)

type I2plog struct {
	verbose bool
}

func (i *I2plog) Error(err error, inp ...interface{}) bool {
	if i.verbose {
		for _, i := range inp {
			os.Stderr.WriteString(i.(string))
		}
	}
	if err != nil {
		log.Println(inp)
		log.Fatal(err)
		return false
	}
	return true
}

func (i *I2plog) Log(inp ...interface{}) bool {
	if i.verbose {
		for _, i := range inp {
			os.Stderr.WriteString(i.(string))
		}
	}
	return true
}
